package main

import (
	"github.com/Zhandos28/social/internal/db"
	"github.com/Zhandos28/social/internal/env"
	"github.com/Zhandos28/social/internal/store"
	"log"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgresql://admin:adminpassword@localhost/postgres?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := store.NewStorage(conn)

	db.Seed(store, conn)
}
