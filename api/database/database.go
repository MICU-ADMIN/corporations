
package database 

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GlobalDB is a global db object that will be used across different packages
var GlobalDB *gorm.DB

// InitDatabase creates a PostgreSQL db connection and stores it in the GlobalDB variable
// It reads the environment variables from the .env file and uses them to create the connection
// It returns an error if the connection fails
func InitDatabase() (err error) {
	// Read the environment variables from the .env file
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error reading .env file")
	}

	// Create the data source name (DSN) using the environment variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config["PGHOST"],
		config["PGUSER"],
		config["PGPASSWORD"],
		config["PGDATABASE"],
		config["PGPORT"],
	)

	// Create the connection and store it in the GlobalDB variable
	GlobalDB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	return
}
