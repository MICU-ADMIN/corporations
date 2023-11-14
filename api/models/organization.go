package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Organization defines the organization in db
// Organization struct is used to store organization information in the database
type Organization struct {
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


// CreateOrganizationRecord creates a organization record in the database
// CreateOrganizationRecord takes a pointer to a Organization struct and creates a organization record in the database
// It returns an error if there is an issue creating the organization record
func (organization *Organization) CreateOrganizationRecord() error {
	result := database.GlobalDB.Create(&organization)

	if result.Error != nil {
    return result.Error
    }

  	to := organization.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on creating your <strong>first organization</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts organization password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (organization *Organization) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  organization.Password = string(bytes)
//  return nil
// }

// CheckPassword checks organization password
// CheckPassword takes a string as a parameter and compares it to the organization's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (organization *Organization) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(organization.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }