package db

import (
	"FinanceApp/config"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

func NewDB(cfg *config.DatabaseConfig) *pgxpool.Pool {
	// Parse the database URL from the configuration
	pgConfig, err := pgxpool.ParseConfig(cfg.URL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	// Configure the connection pool settings
	pgConfig.MaxConns = int32(cfg.MaxConns)
	pgConfig.MinConns = int32(cfg.MinConns)
	pgConfig.HealthCheckPeriod = 10 * time.Second
	pgConfig.MaxConnLifetime = 30 * time.Minute
	pgConfig.MaxConnIdleTime = 5 * time.Minute
	pgConfig.ConnConfig.ConnectTimeout = 5 * time.Second

	// Create the connection pool with the configured settings
	dbpool, err := pgxpool.NewWithConfig(context.Background(), pgConfig)
	if err != nil {
		log.Fatalf("Unable to create database connection pool: %v", err)
	}

	return dbpool
}

func CloseDB(dbpool *pgxpool.Pool) {
	dbpool.Close()
}
