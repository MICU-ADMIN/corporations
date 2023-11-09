package hadith_service

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var admin_db *sql.DB
var err error


func InitializeDB_Hadith() (*sql.DB, error) {

	err := godotenv.Load()
  if err != nil {
    log.Println("Error loading .env file")
  }

	// create mysql connection 

	// Open a connection to the database
	admin_db, err = sql.Open("mysql", os.Getenv("ADMINMYSQL_DSN"))
	if err != nil {
		log.Fatal("failed to open admin_db connection", err)
	}

// // Ping the database
if err = admin_db.Ping(); err != nil {
    log.Fatal("failed to ping database", err)
 }

    return admin_db, nil
}

type Hadith struct {
	Id             int
	CollectionId   int
	BookId         int
	HadithNumber   int
	Label          string
	Arabic         string
	EnglishTrans   string
	PrimaryNarrator string
}



func GetHadith(c *gin.Context) {
	query := "SELECT * FROM hadith"
db, err := InitializeDB_Hadith()
	if err != nil {
		log.Fatal("(GetHadiths) InitializeDB_Hadith ", err)
	}
res, err := db.Query(query)
	if err != nil {
		log.Fatal("(GetHadiths) DB_Hadith.Query", err)
	}
defer res.Close()
	if err != nil {
		log.Fatal("(GetHadiths) DB_Hadith.Res", err)
	}

	hadiths := []Hadith{}
	for res.Next() {
		var hadith Hadith
        err := res.Scan(&hadith.Id, &hadith.CollectionId, &hadith.BookId, &hadith.HadithNumber, &hadith.Label, &hadith.Arabic, &hadith.EnglishTrans, &hadith.PrimaryNarrator)
		if err != nil {
			log.Fatal("(GetHadiths) res.Scan", err)
		}
		hadiths = append(hadiths, hadith)
	}
db.Close()
	c.JSON(http.StatusOK, hadiths)
}

func GetAhadith(c *gin.Context) {
	query := "SELECT * FROM hadith"
db, err := InitializeDB_Hadith()
	if err != nil {
		log.Fatal("(GetHadiths) InitializeDB_Hadith ", err)
	}
res, err := db.Query(query)
	if err != nil {
		log.Fatal("(GetHadiths) DB_Hadith.Query", err)
	}
defer res.Close()
	if err != nil {
		log.Fatal("(GetHadiths) DB_Hadith.Res", err)
	}

	hadiths := []Hadith{}
	for res.Next() {
		var hadith Hadith
        err := res.Scan(&hadith.Id, &hadith.CollectionId, &hadith.BookId, &hadith.HadithNumber, &hadith.Label, &hadith.Arabic, &hadith.EnglishTrans, &hadith.PrimaryNarrator)
		if err != nil {
			log.Fatal("(GetHadiths) res.Scan", err)
		}
		hadiths = append(hadiths, hadith)
	}
db.Close()
	c.JSON(http.StatusOK, hadiths)
}

func GetAhadithFiltered(c *gin.Context) {
    // Get collectionId and bookId from query parameters
    collectionID := c.Query("collectionId")
    bookID := c.Query("bookId")

    // Check if either collectionId or bookId is provided
    if collectionID == "" || bookID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Any of collectionId, bookId may be  missing.",
        })
        return
    }

    // Construct the SQL query
    query := "SELECT id, collectionId, bookId, hadithNumber, label, arabic, englishTrans, primaryNarrator FROM hadith WHERE 1"

    // Add conditions based on provided parameters
    if collectionID != "" {
        query += " AND collectionId = " + collectionID
    }
    if bookID != "" {
        query += " AND bookId = " + bookID
    }

	db, err := InitializeDB_Hadith()
	if err != nil {
		log.Fatal("(GetAhadithFiltered) InitializeDB_Hadith ", err)
	}
res, err := db.Query(query)
	if err != nil {
            log.Println("query",query)
		log.Fatal("(GetAhadithFiltered) GetAhadithFiltered.Query", err)
	}
defer res.Close()
	if err != nil {
		log.Fatal("(GetAhadithFiltered) GetAhadithFiltered.Res", err)
	}

	hadiths := []Hadith{}
	for res.Next() {
		var hadith Hadith
        err := res.Scan(&hadith.Id, &hadith.CollectionId, &hadith.BookId, &hadith.HadithNumber, &hadith.Label, &hadith.Arabic, &hadith.EnglishTrans, &hadith.PrimaryNarrator)
		if err != nil {
			log.Fatal("(GetAhadithFiltered) res.Scan", err)
		}
		hadiths = append(hadiths, hadith)
	}
db.Close()
	c.JSON(http.StatusOK, hadiths)
}


// func GetSingleHadith(c *gin.Context) {
// 	hadithId := c.Param("hadithId")
// 	hadithId = strings.ReplaceAll(hadithId, "/", "")
// 	hadithIdInt, err := strconv.Atoi(hadithId)
// 	if err != nil {
// 		log.Fatal("(GetSingleHadith) strconv.Atoi", err)
// 	}

// 	var hadith Hadith
// 	query := `SELECT * FROM hadiths WHERE id = ?`
// 	err = InitializeDB_Hadith.QueryRow(query, hadithIdInt).Scan(&hadith.Id, &hadith.Name, &hadith.Price)
// 	if err != nil {
// 		log.Fatal("(GetSingleHadith) InitializeDB_Hadith.Exec", err)
// 	}

// 	c.JSON(http.StatusOK, hadith)
// }

// func CreateHadith(c *gin.Context) {
// 	var newHadith Hadith
// 	err := c.BindJSON(&newHadith)
// 	if err != nil {
// 		log.Fatal("(CreateHadith) c.BindJSON", err)
// 	}

// 	query := `INSERT INTO Hadith (name, price) VALUES (?, ?)`
// 	res, err := InitializeDB_Hadith.Exec(query, newHadith.Name, newHadith.Price)
// 	if err != nil {
// 		log.Fatal("(CreateHadith) InitializeDB_Hadith.Exec", err)
// 	}
// 	newHadith.Id, err = res.LastInsertId()
// 	if err != nil {
// 		log.Fatal("(CreateHadith) res.LastInsertId", err)
// 	}

// 	c.JSON(http.StatusOK, newHadith)
// }

// func UpdateHadith(c *gin.Context) {
// 	var updates Hadith
// 	err := c.BindJSON(&updates)
// 	if err != nil {
// 		log.Fatal("(UpdateHadith) c.BindJSON", err)
// 	}

// 	hadithId := c.Param("hadithId")
// 	hadithId = strings.ReplaceAll(hadithId, "/", "")
// 	hadithIdInt, err := strconv.Atoi(hadithId)
// 	if err != nil {
// 		log.Fatal("(UpdateHadith) strconv.Atoi", err)
// 	}

// 	query := `UPDATE Hadith SET name = ?, price = ? WHERE id = ?`
// 	_, err = InitializeDB_Hadith.Exec(query, updates.Name, updates.Price, hadithIdInt)
// 	if err != nil {
// 		log.Fatal("(UpdateHadith) InitializeDB_Hadith.Exec", err)
// 	}

// 	c.Status(http.StatusOK)
// }

// func DeleteHadith(c *gin.Context) {
// 	hadithId := c.Param("hadithId")

// 	hadithId = strings.ReplaceAll(hadithId, "/", "")
// 	hadithIdInt, err := strconv.Atoi(hadithId)
// 	if err != nil {
// 		log.Fatal("(DeleteHadith) strconv.Atoi", err)
// 	}
// 	query := `DELETE FROM Hadith WHERE id = ?`
// 	_, err = InitializeDB_Hadith.Exec(query, hadithIdInt)
// 	if err != nil {
// 		log.Fatal("(DeleteHadith) InitializeDB_Hadith.Exec", err)
// 	}

// 	c.Status(http.StatusOK)
// }
