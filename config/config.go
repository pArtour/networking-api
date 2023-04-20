package config

import (
	"os"
)

func GetDBConnectionString() string {
	// Replace with your method of fetching the connection string
	return os.Getenv("DB_CONNECTION_STRING")
}
