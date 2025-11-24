package main

import (
	"log"

	in "github.com/t4e1/clean-go/clean-b/post/adapter/in/http"
	"github.com/t4e1/clean-go/clean-b/post/adapter/out"
	"github.com/t4e1/clean-go/clean-b/post/configs/setup"
	"github.com/t4e1/clean-go/clean-b/post/internal/core/usecase"
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

	// Initiate Outbound Adapter
	out := out.InitPostgresAdapter(db)

	// Initiate Usecase
	uc := usecase.InitRESTUsecase(out)

	// Initiate Inbound Adapter
	in := in.InitRESTAdapter(uc)
}
