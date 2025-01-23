package main

import (
	"context"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	baseStorage "github.com/khrulsergey/chain-processor/pkg/storage"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"listener/http/rest"
	"listener/http/router"
	"listener/service"
	"listener/storage"
	"listener/web3/connector"
	"listener/web3/reader"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	appConfig   *config.AppConfig
	dbConfig    *config.PgDBConfig
	appLogger   logger.Logger
	dbConnector *gorm.DB

	//Service
	transactionService service.TransactionService
	kafkaService       service.KafkaService

	//REST
	versionController *rest.SystemController
)

type shutdown = func() error

func main() {
	log.Println("INFO: start initialization")

	if err := initializeDependencies(); err != nil {
		log.Fatalf("FATAL: initialization failed: %v", err)
	}

	log.Println("INFO: initialization finished")

	errCh := make(chan error, 1)
	httpServerShutdown, err := startHTTPServer(errCh)
	if err != nil {
		appLogger.Fatalf("Server failed to start HTTP: %v", err)
	}

	appLogger.Info("Server ready to accept connections")

	waitServers(
		errCh,
		httpServerShutdown,
	)
}

func initializeDependencies() error {
	var err error

	ctx := context.Background()
	appLogger, err = logger.NewRelease()
	defer func() {
		_ = appLogger.Flush()
	}() // flushes buffer, if any

	// Define new appConfig
	appConfig, err = config.InitAppConfig()
	if err != nil {
		appLogger.Fatalf("Unable to read configuration %v", err)
	}

	//update logger level
	if appConfig.LoggerMode == "DEBUG" {
		appLogger, _ = logger.NewDebug()
	}

	//Define database connector
	{
		dbConfig, err = config.NewPgDBConfig()
		if err != nil {
			appLogger.Fatalf("Unable to read database configuration: %v", err)
		}

		dbConnector, err = baseStorage.NewGormPgDB(dbConfig)
		if err != nil {
			appLogger.Fatalf("Unable to establish database: %v", err)
		}

		appLogger.Info("applying db migration")
		if err := storage.MigrateGorm(dbConnector); err != nil {
			log.Fatalf("db migration failed %v", err)
		}
	}

	//Define transaction services
	{
		transactionStorage := storage.NewTransactionStorageGorm(dbConnector)
		transactionService = service.NewTransactionService(transactionStorage, appLogger)
	}

	chainConfigs, err := config.NewChainConfigs()
	if err != nil {
		appLogger.Fatalf("Unable to read chain configuration %v", err)
	}

	appLogger.Info("started app with chain config enabled for chains:", len(chainConfigs.GetAllEnabledChains()))

	//Define kafka services
	{
		kafkaConfig, err := config.NewKafkaConfig()
		if err != nil {
			appLogger.Fatalf("Unable to read kafka configuration: %v", err)
		}
		kafkaService = service.NewKafkaService(kafkaConfig, appLogger)
	}

	//Define chain services
	{
		eventProcessor := service.NewEventProcessor(transactionService, kafkaService, appLogger)
		connectionManager := connector.NewChainConnectorManager(chainConfigs, appLogger)
		appLogger.Info("starting chain connector manager with enabled chain size: ",
			connectionManager.GetChainConnectorSize())

		for _, cConfig := range chainConfigs.GetAllEnabledChains() {
			_ = reader.NewOneInchContractListener(ctx, cConfig, connectionManager, eventProcessor,
				transactionService, appLogger)
		}
	}

	//Define main Controllers
	versionController = rest.NewSystemController(appConfig, appLogger)

	return err
}

func startHTTPServer(errCh chan<- error) (shutdown, error) {
	httpRouter := router.NewRouterHandler(appLogger, *appConfig, versionController)
	appRouter := httpRouter.InitRouter()

	//init HTTP httpServer
	httpServer := &fasthttp.Server{
		Handler: appRouter.Handler,
	}

	listeningAddress := net.JoinHostPort(appConfig.Host, appConfig.Port)
	log.Printf("INFO: starting HTTP httpServer listening %v\n", listeningAddress)
	go func() {
		errCh <- httpServer.ListenAndServe(net.JoinHostPort(appConfig.Host, appConfig.Port))
	}()

	return httpServer.Shutdown, nil
}

func waitServers(errCh <-chan error, shutdowns ...shutdown) {
	done := make(chan struct{})
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		select {
		case err := <-errCh:
			log.Printf("ERROR: got an error: %v\n", err)
		case s := <-signalCh:
			log.Printf("INFO: got a signal: %v\n", s)
		}
		log.Println("INFO: gracefully shutting down...")
		for _, sd := range shutdowns {
			if err := sd(); err != nil {
				log.Printf("ERROR: error during shutdown: %v\n", err)
			}
		}
		close(done)
	}()
	<-done
}
