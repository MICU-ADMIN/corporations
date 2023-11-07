package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Hadith struct {
	Id    int64
	Name  string
	Price int
}

func main() {

err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}


// Ping the database
if err = db.Ping(); err != nil {
    log.Fatal("failed to ping database", err)
 }


	router := gin.Default()

router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

// hadiths
	router.GET("/hadiths", GetHadiths)


	
	router.Run()
}

