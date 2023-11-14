```go


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


```

```mermaid

Here is a Mermaid Markdown overview for the Go file you provided:
```mermaid
graph LR
    A[GlobalDB] --> B[InitDatabase]
    B --> C[Read environment variables from .env file]
    C --> D[Create DSN using environment variables]
    D --> E[Create PostgreSQL connection using DSN]
    E --> F[Store connection in GlobalDB]
    F --> G[Return error if connection fails]
    G --> H[End]
```
In this overview, the `GlobalDB` node represents the global database object, which is the main entry point for the code. The `InitDatabase` function is called to create a PostgreSQL connection and store it in the `GlobalDB` variable.
The `InitDatabase` function is divided into several steps:
1. `Read environment variables from .env file`: This step reads the environment variables from the `.env` file using the `godotenv` package.
2. `Create DSN using environment variables`: This step creates the data source name (DSN) using the environment variables read in step 1.
3. `Create PostgreSQL

```
