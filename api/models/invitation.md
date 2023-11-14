```go

package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Invitation defines the invitation in db
// Invitation struct is used to store invitation information in the database
type Invitation struct {
 gorm.Model
 ID       int    `gorm:"primaryKey"`
 Name     string `json:"name" binding:"required"`
 Email    string `json:"email" binding:"required" gorm:"unique"`
 Password string `json:"password" binding:"required"`
}

// CreateInvitationRecord creates a invitation record in the database
// CreateInvitationRecord takes a pointer to a Invitation struct and creates a invitation record in the database
// It returns an error if there is an issue creating the invitation record
func (invitation *Invitation) CreateInvitationRecord() error {
	result := database.GlobalDB.Create(&invitation)

	if result.Error != nil {
    return result.Error
    }

  	to := invitation.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on sending your <strong>first email</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts invitation password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func (invitation *Invitation) HashPassword(password string) error {
 bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
 if err != nil {
  return err
 }
 invitation.Password = string(bytes)
 return nil
}

// CheckPassword checks invitation password
// CheckPassword takes a string as a parameter and compares it to the invitation's encrypted password
// It returns an error if there is an issue comparing the passwords
func (invitation *Invitation) CheckPassword(providedPassword string) error {
 err := bcrypt.CompareHashAndPassword([]byte(invitation.Password), []byte(providedPassword))
 if err != nil {
  return err
 }
 return nil
}

```

```mermaid

Here is a Mermaid Markdown overview for the Go file you provided:
```mermaid
graph LR
    Invitation[Invitation] --> Database[database]
    Invitation --> CreateInvitationRecord[CreateInvitationRecord]
    Invitation --> HashPassword[HashPassword]
    Invitation --> CheckPassword[CheckPassword]
    Invitation --> SendEmail[SendEmail]
```
Explanation:

* `Invitation` is the main struct in the file, representing an invitation.
* `Database` is the database connection, which is used to store and retrieve invitation records.
* `CreateInvitationRecord` is a function that creates a new invitation record in the database. It takes a pointer to an `Invitation` struct as an argument and returns an error if there is an issue creating the record.
* `HashPassword` is a function that encrypts the invitation's password using the `bcrypt` package. It takes a string as an argument and returns an error if there is an issue encrypting the password.
* `CheckPassword` is a function that compares the invitation's encrypted password to a provided password.

```
