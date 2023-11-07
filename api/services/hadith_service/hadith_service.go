package hadith_service

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func InitializeDB_Hadith() (*sql.DB, error) {
 err := godotenv.Load()
  if err != nil {
    log.Println("Error loading .env file")
  }
// not crucial only needed in dev as our providers have their own methods for injecting env 

	// Get environment variables
    host := os.Getenv("HADITHPOSTGRES_HOST")
		if host == "" {
		log.Fatal("InitializeDB_Hadith", "(host is empty)")
		}
    port := os.Getenv("HADITHPOSTGRES_PORT")
	if port == "" {
		log.Fatal("InitializeDB_Hadith", "(port is empty)")
		}
    user := os.Getenv("HADITHPOSTGRES_USER")
	if user == "" {
		log.Fatal("InitializeDB_Hadith", "(user is empty)")
		}
    password := os.Getenv("HADITHPOSTGRES_PASSWORD")
	if password == "" {
		log.Fatal("InitializeDB_Hadith", "(password is empty)")
		}
    dbname := os.Getenv("HADITHPOSTGRES_NAME")
	if dbname == "" {
		log.Fatal("InitializeDB_Hadith", "dbname is empty")
		}

    // Construct the PostgreSQL connection string
    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)


    // Open a database connection
    hadithDB, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Test the database connection
    err = hadithDB.Ping()
    if err != nil {
        return nil, err
    }

    return hadithDB, nil
}



type Hadith struct {
	Id    int64
	Name  string
	Price int
}



func GetHadiths(c *gin.Context) {
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
		err := res.Scan(&hadith.Id, &hadith.Name, &hadith.Price)
		if err != nil {
			log.Fatal("(GetHadiths) res.Scan", err)
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
