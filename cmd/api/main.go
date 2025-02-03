package main

import (
	"github.com/Zhandos28/social/internal/db"
	"github.com/Zhandos28/social/internal/env"
	"github.com/Zhandos28/social/internal/store"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"log"
)

func main() {
	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	err := godotenv.Load()
	if err != nil {
		logger.Fatalw("Error loading .env file", "error", err)
	}

	cfg := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:           env.GetString("DB_ADDR", "postgres//admin:adminpassword@localhost:5432/social?sslmode=disable"),
			maxOpenConns:   env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:   env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTimeout: env.GetString("DB_MAX_IDLE_TIMEOUT", "15m"),
		},
		env:     env.GetString("ENV", "development"),
		version: env.GetString("VERSION", "0.0.1"),
	}

	db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTimeout)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)

	app := application{
		config: cfg,
		store:  store,
		logger: logger,
	}

	mux := app.mount()
	logger.Fatal(app.run(mux))
}
