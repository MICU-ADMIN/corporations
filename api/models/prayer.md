package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Prayer defines the prayer in db
// Prayer struct is used to store prayer information in the database
type Prayer struct {
    ID          int       `gorm:"primaryKey" json:"id"`
    OwnerID     string       
    Name        string    `json:"name" binding:"required"`
    Email       string    `json:"email" binding:"required" gorm:"unique"`
    Phone       string    `json:"phone" binding:"required"`
    Description string    `json:"description"`
    Services    []string    `json:"services"`
    Images      []string  `json:"images"`
    // Add other fields as needed
}


// CreatePrayerRecord creates a prayer record in the database
// CreatePrayerRecord takes a pointer to a Prayer struct and creates a prayer record in the database
// It returns an error if there is an issue creating the prayer record
func (prayer *Prayer) CreatePrayerRecord() error {
	result := database.GlobalDB.Create(&prayer)

	if result.Error != nil {
    return result.Error
    }

  	to := prayer.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on creating your <strong>first prayer</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts prayer password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (prayer *Prayer) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  prayer.Password = string(bytes)
//  return nil
// }

// CheckPassword checks prayer password
// CheckPassword takes a string as a parameter and compares it to the prayer's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (prayer *Prayer) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(prayer.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }