package main

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)


var db *sql.DB

type Hadith struct {
	Id    int64
	Name  string
	Price int
}

func main() {
	// Load in the `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("failed to load env", err)
	}

	// Open a connection to the database
	db, err = sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatal("failed to open db connection", err)
	}

	// Build router & define routes


	router := gin.Default()
// add swagger
router.GET ("/docs/*any", ginSwagger .WrapHandler(swaggerFiles.Handler))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	router.GET("/hadiths", GetHadiths)
	router.GET("/hadiths/:hadithId", GetSingleHadith)
	router.POST("/hadiths", CreateHadith)
	
// curl -X POST http://localhost:8080/hadiths \
// -H "Content-Type: application/json" \
// -d '{
//   "name": "YourHadithName",
//   "price": 123
// }'

	router.PUT("/hadiths/:hadithId", UpdateHadith)
	router.DELETE("/hadiths/:hadithId", DeleteHadith)

	// Run the router
	router.Run()
}

func GetHadiths(c *gin.Context) {
	query := "SELECT * FROM Hadith"
	res, err := db.Query(query)
	defer res.Close()
	if err != nil {
		log.Fatal("(GetHadiths) db.Query", err)
	}

	hadiths := []Hadith{}
	for res.Next() {
		var hadith Hadith
		err := res.Scan(&hadith.Id, &hadith.Name, &hadith.Price)
		if err != nil {
			log.Fatal("(GetHadiths) res.Scan", err)
		}
		hadiths = append(hadiths, hadith)
	}

	c.JSON(http.StatusOK, hadiths)
}

func GetSingleHadith(c *gin.Context) {
	hadithId := c.Param("hadithId")
	hadithId = strings.ReplaceAll(hadithId, "/", "")
	hadithIdInt, err := strconv.Atoi(hadithId)
	if err != nil {
		log.Fatal("(GetSingleHadith) strconv.Atoi", err)
	}

	var hadith Hadith
	query := `SELECT * FROM hadiths WHERE id = ?`
	err = db.QueryRow(query, hadithIdInt).Scan(&hadith.Id, &hadith.Name, &hadith.Price)
	if err != nil {
		log.Fatal("(GetSingleHadith) db.Exec", err)
	}

	c.JSON(http.StatusOK, hadith)
}

func CreateHadith(c *gin.Context) {
	var newHadith Hadith
	err := c.BindJSON(&newHadith)
	if err != nil {
		log.Fatal("(CreateHadith) c.BindJSON", err)
	}

	query := `INSERT INTO Hadith (name, price) VALUES (?, ?)`
	res, err := db.Exec(query, newHadith.Name, newHadith.Price)
	if err != nil {
		log.Fatal("(CreateHadith) db.Exec", err)
	}
	newHadith.Id, err = res.LastInsertId()
	if err != nil {
		log.Fatal("(CreateHadith) res.LastInsertId", err)
	}

	c.JSON(http.StatusOK, newHadith)
}

func UpdateHadith(c *gin.Context) {
	var updates Hadith
	err := c.BindJSON(&updates)
	if err != nil {
		log.Fatal("(UpdateHadith) c.BindJSON", err)
	}

	hadithId := c.Param("hadithId")
	hadithId = strings.ReplaceAll(hadithId, "/", "")
	hadithIdInt, err := strconv.Atoi(hadithId)
	if err != nil {
		log.Fatal("(UpdateHadith) strconv.Atoi", err)
	}

	query := `UPDATE Hadith SET name = ?, price = ? WHERE id = ?`
	_, err = db.Exec(query, updates.Name, updates.Price, hadithIdInt)
	if err != nil {
		log.Fatal("(UpdateHadith) db.Exec", err)
	}

	c.Status(http.StatusOK)
}

func DeleteHadith(c *gin.Context) {
	hadithId := c.Param("hadithId")

	hadithId = strings.ReplaceAll(hadithId, "/", "")
	hadithIdInt, err := strconv.Atoi(hadithId)
	if err != nil {
		log.Fatal("(DeleteHadith) strconv.Atoi", err)
	}
	query := `DELETE FROM Hadith WHERE id = ?`
	_, err = db.Exec(query, hadithIdInt)
	if err != nil {
		log.Fatal("(DeleteHadith) db.Exec", err)
	}

	c.Status(http.StatusOK)
}
