package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

type Hadith struct {
	Id    int64
	Name  string
	Price int
}

func main() {

db, err := sql.Open("mysql", "1m8vf791734q5kf3rcjf:pscale_pw_JEJLP3rdncDmYOfFoZWWftODzvJEFK8IoAvNiTXw6HA@tcp(aws.connect.psdb.cloud)/backend?tls=true&interpolateParams=true",
  )
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


// services 


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
