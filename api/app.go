package main

import (
	"database/sql"
	"net/http"

	"mosque.icu.corporations.api/services/hadith_service"

	"github.com/gin-gonic/gin"
)

//  gobal variables
var admin_db *sql.DB


func main() {

	

	// create mysql connection 

	// Open a connection to the database
// 	admin_db, err = sql.Open("mysql", os.Getenv("DSN"))
// 	if err != nil {
// 		log.Fatal("failed to open admin_db connection", err)
// 	}


// // Ping the database
// if err = admin_db.Ping(); err != nil {
//     log.Fatal("failed to ping database", err)
//  }


// create the router 


	router := gin.Default()

router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})

// hadiths
	router.GET("/hadiths", hadith_service.GetHadiths)

	router.Run()

	
}


