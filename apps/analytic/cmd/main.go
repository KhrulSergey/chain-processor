package main

import (
	"analytic/http/rest"
	"analytic/http/router"
	"analytic/service"
	"analytic/storage"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	baseStorage "github.com/khrulsergey/chain-processor/pkg/storage"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	appConfig   *config.AppConfig
	dbConfig    *config.ChDBConfig
	appLogger   logger.Logger
	dbConnector *gorm.DB

	////Storage
	//complianceRepository storage.ComplianceRepository

	//Service
	dataProcessor service.DataProcessor
	kafkaService  service.KafkaService
	//externalComplianceService    service.ExternalComplianceService
	//accountOperatorServiceClient service.AccountOperatorServiceClient
	//complianceService            service.CompliancePersonService

	//REST
	systemController *rest.SystemController
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
		log.Fatalf("FATAL: failed to start HTTP server: %v", err)
	}

	log.Println("INFO: server ready to accept connections")

	waitServers(
		errCh,
		httpServerShutdown,
	)
}

func initializeDependencies() error {
	var err error

	//ctx := context.Background()
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
		dbConfig, err = config.NewChDBConfig()
		if err != nil {
			appLogger.Fatalf("Unable to read database configuration: %v", err)
		}

		dbConnector, err = baseStorage.NewGormChDB(dbConfig)
		if err != nil {
			appLogger.Fatalf("Unable to establish database: %v", err)
		}

		appLogger.Info("applying db migration")
		if err := storage.MigrateGorm(dbConnector); err != nil {
			log.Fatalf("db migration failed %v", err)
		}
	}

	// Define Repository
	tradeStorage := storage.NewTradeStorageGorm(dbConnector, appLogger)

	// Define main services
	{
		tradeService := service.NewTradeService(tradeStorage, appLogger)
		dataProcessor = service.NewDataProcessor(tradeService, appLogger)
	}

	//Define kafka services
	{
		kafkaConfig, err := config.NewKafkaConfig()
		if err != nil {
			appLogger.Fatalf("Unable to read kafka configuration: %v", err)
		}
		kafkaService = service.NewKafkaService(kafkaConfig, dataProcessor, appLogger)
	}

	////Define service for interact with different compliance providers
	//externalComplianceService = service.InitExternalComplianceService(appLogger)
	//
	////Define service-client for interact with AO-service
	//accountOperatorServiceClient = service.InitAccountOperatorServiceClient(appLogger)
	//

	//complianceService = service.InitComplianceService(appLogger, complianceRepository,
	//	externalComplianceService, accountOperatorServiceClient)

	//todo delete all REST layer. Use only gRPC
	//Define main Controllers
	systemController = rest.NewSystemController(appConfig, appLogger)

	return err
}

func startHTTPServer(errCh chan<- error) (shutdown, error) {
	httpRouter := router.NewRouterHandler(appLogger, *appConfig, systemController)
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
