package main

import (
	"log"

	"github.com/t4e1/clean-go/clean-b/post/configs/setup"
	"gorm.io/gorm"
)

// Cmd package for using server's entry point.
var db *gorm.DB

func main() {
	// Initiate database connection
	db, err := setup.InitPostgres()
	if err != nil {
		log.Fatalf("Failed to initialize postgresql connection: %v", err)
	}

	// Initiate handler(adpater/in)

}
