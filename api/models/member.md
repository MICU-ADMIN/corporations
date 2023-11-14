```go

package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Member defines the member in db
// Member struct is used to store member information in the database
type Member struct {
    ID          int       `gorm:"primaryKey" json:"id"`
    OrganizationID    string    `json:"organizationID"  binding:"required"`   
    Name        string    `json:"name"`
    Email       string    `json:"email"  binding:"required" gorm:"unique"`
    Phone       string    `json:"phone"`
    Bio      string    `json:"bio"`
	Role    string    `json:"role"`
    Photo      string  `json:"photo"`
    // Add other fields as needed
}


// CreateMemberRecord creates a member record in the database
// CreateMemberRecord takes a pointer to a Member struct and creates a member record in the database
// It returns an error if there is an issue creating the member record
func (member *Member) CreateMemberRecord() error {
	result := database.GlobalDB.Create(&member)

	if result.Error != nil {
    return result.Error
    }

  	to := member.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Welcome to the  <strong>"+  member.OrganizationID +" organization</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts member password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (member *Member) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  member.Password = string(bytes)
//  return nil
// }

// CheckPassword checks member password
// CheckPassword takes a string as a parameter and compares it to the member's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (member *Member) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(member.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }

```
