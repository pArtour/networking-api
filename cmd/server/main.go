package main

import (
	"github.com/pArtour/networking-api/config"
	"github.com/pArtour/networking-api/internal/handlers"
	"github.com/pArtour/networking-api/internal/storage"

	"log"
	"net/http"
	"os"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Define your routes here
	//router.HandleFunc("/users/register", registerUserHandler).Methods(http.MethodPost)
	//router.HandleFunc("/users/login", loginUserHandler).Methods(http.MethodPost)
	//router.HandleFunc("/users/{id}", getUserHandler).Methods(http.MethodGet)
	router.HandleFunc("/example", handlers.GetExampleHandler).Methods(http.MethodGet)
	// Add more routes as needed

	return router
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT not set in .env file")
	}

	dbConnectionString := config.GetDBConnectionString()
	if err := storage.ConnectDB(dbConnectionString); err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	defer storage.CloseDB()

	router := NewRouter()

	log.Printf("Server listening on port %s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, router))
}
