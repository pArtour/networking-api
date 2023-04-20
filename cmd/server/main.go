package main

import (
	"os"
	"log"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/pArtour/networking-api/config"
	"github.com/pArtour/networking-api/internal/storage"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err)
	}

	dbConnectionString := config.GetDBConnectionString()
	if err := storage.ConnectDB(dbConnectionString); err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer storage.CloseDB()

	api := operations.NewYourAppNameAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()

	// Implement your API handlers here
	api.GetExampleHandler = handlers.GetExampleHandler()

	// Set the server address and port
	server.Host = "localhost"
	server.Port = 8080

	// Start the server
	log.Printf("Starting server on %s:%d", server.Host, server.Port)
	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}