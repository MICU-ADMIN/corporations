```go

package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Display defines the display in db
// Display struct is used to store display information in the database
type Display struct {
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


// CreateDisplayRecord creates a display record in the database
// CreateDisplayRecord takes a pointer to a Display struct and creates a display record in the database
// It returns an error if there is an issue creating the display record
func (display *Display) CreateDisplayRecord() error {
	result := database.GlobalDB.Create(&display)

	if result.Error != nil {
    return result.Error
    }

  	to := display.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on creating your <strong>first display</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts display password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (display *Display) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  display.Password = string(bytes)
//  return nil
// }

// CheckPassword checks display password
// CheckPassword takes a string as a parameter and compares it to the display's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (display *Display) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(display.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }

```


```mermaid

graph LR
    A[CreateDisplayRecord] --> B[database.GlobalDB.Create(&display)]
    B --> C[result.Error]
    C --> D[if result.Error != nil]
    D --> E[return result.Error]

    A --> F[to] --> G[from] --> H[subject] --> I[html]
    F --> K[lib.SendEmail(to, from, subject, html)]
    K --> L[if err != nil] --> M[log.Println("Error sending email:", err)]
    L --> M[if success] --> N[log.Println("Email sent successfully.")]

    A --> J[HashPassword] --> K[bcrypt.GenerateFromPassword([]byte(password), 14)]
    K --> L[if err != nil] --> M[log.Println("Error encrypting password:", err)]
    L --> M[display.Password = string(bytes)] --> N[return nil]

    A --> P[

```
