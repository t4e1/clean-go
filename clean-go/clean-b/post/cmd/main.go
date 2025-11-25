package main

import (
	"fmt"
	"log"

	"github.com/t4e1/clean-go/clean-b/post/adapter/in"
	"github.com/t4e1/clean-go/clean-b/post/adapter/out"
	"github.com/t4e1/clean-go/clean-b/post/config/setup"
	"github.com/t4e1/clean-go/clean-b/post/internal/core/usecase"
	"gorm.io/gorm"
)

// Cmd package for using server's entry point.
var db *gorm.DB

const (
	portNum = 10001
)

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

	// Initiate Gin Router
	r := setup.SetupRouter(in)

	addr := fmt.Sprintf(":%d", portNum)
	r.Run(addr)
}
