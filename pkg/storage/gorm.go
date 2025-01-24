package storage

import (
	"crypto/tls"
	"fmt"
	stdck "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/khrulsergey/chain-processor/config"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewGormPgDB(config *config.PgDBConfig) (*gorm.DB, error) {
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s %s",
		config.DatabaseHost,
		config.DatabasePort,
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseName,
		sslConnectionParams(config.DatabaseRootCA),
	)
	conn, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: connectionString,
			},
		),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}
	db, err := conn.DB()
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Second * time.Duration(config.ConnMaxLifeTime))
	return conn, err
}

func sslConnectionParams(dbRootCA string) string {
	if dbRootCA == "" {
		return "sslmode=disable"
	}
	return fmt.Sprintf("sslmode=verify-full sslrootcert=%s", dbRootCA)
}

func NewGormChDB(config *config.ChDBConfig) (*gorm.DB, error) {
	var chTls *tls.Config = nil
	if config.DatabaseSSL {
		chTls = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	sqlDB := stdck.OpenDB(&stdck.Options{
		Protocol: stdck.HTTP,
		Addr:     []string{config.DatabaseHost + ":" + config.DatabasePort},
		Auth: stdck.Auth{
			Database: config.DatabaseName,
			Username: config.DatabaseUser,
			Password: config.DatabasePassword,
		},
		TLS: chTls,
		Settings: stdck.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &stdck.Compression{
			Method: stdck.CompressionLZ4,
		},
		Debug: true,
	})

	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
		Conn: sqlDB, // initialize with existing database conn
	}))
	return db, err
}

type User struct {
	Name string
	Age  int
}
