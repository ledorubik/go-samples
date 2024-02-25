package main

import (
	"go-samples/2-init-db-gorm-postgres/config"
	_db "go-samples/2-init-db-gorm-postgres/db"
	"log"
)

func main() {
	// Init config
	cfg := config.New()

	// Init database
	db, err := _db.InitDB(cfg)
	if err != nil {
		log.Fatalf("database init error: %s", err)
	}

	log.Printf("db: %v", db)
}
