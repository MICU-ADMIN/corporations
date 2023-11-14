package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Event defines the event in db
// Event struct is used to store event information in the database
type Event struct {
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


// CreateEventRecord creates a event record in the database
// CreateEventRecord takes a pointer to a Event struct and creates a event record in the database
// It returns an error if there is an issue creating the event record
func (event *Event) CreateEventRecord() error {
	result := database.GlobalDB.Create(&event)

	if result.Error != nil {
    return result.Error
    }

  	to := event.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on creating your <strong>first event</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts event password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (event *Event) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  event.Password = string(bytes)
//  return nil
// }

// CheckPassword checks event password
// CheckPassword takes a string as a parameter and compares it to the event's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (event *Event) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(event.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }