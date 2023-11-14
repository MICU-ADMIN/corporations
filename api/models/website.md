```go

package models

import (
	"go-jwt/database"
	"go-jwt/lib"
	"log"
)

// Website defines the website in db
// Website struct is used to store website information in the database
type Website struct {
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


// CreateWebsiteRecord creates a website record in the database
// CreateWebsiteRecord takes a pointer to a Website struct and creates a website record in the database
// It returns an error if there is an issue creating the website record
func (website *Website) CreateWebsiteRecord() error {
	result := database.GlobalDB.Create(&website)

	if result.Error != nil {
    return result.Error
    }

  	to := website.Email
    from := "noreply@mosque.icu"
    subject := "Hello World"
    html := "<p>Congrats on creating your <strong>first website</strong>!</p>"

    success, err := lib.SendEmail(to, from, subject, html)
    if err != nil {
        log.Println("Error sending email:", err)
        
    }

    if success {
        log.Println("Email sent successfully.")
    }

 return nil
}

// HashPassword encrypts website password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
// func (website *Website) HashPassword(password string) error {
//  bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//  if err != nil {
//   return err
//  }
//  website.Password = string(bytes)
//  return nil
// }

// CheckPassword checks website password
// CheckPassword takes a string as a parameter and compares it to the website's encrypted password
// It returns an error if there is an issue comparing the passwords
// func (website *Website) CheckPassword(providedPassword string) error {
//  err := bcrypt.CompareHashAndPassword([]byte(website.Password), []byte(providedPassword))
//  if err != nil {
//   return err
//  }
//  return nil
// }

```

```mermaid

Here is a Mermaid Markdown overview for the Go file you provided:
```mermaid
graph LR
    A[Website] --> B[Database]
    A --> C[CreateWebsiteRecord]
    A --> D[HashPassword]
    A --> E[CheckPassword]
    B --> F[Save Website Record]
    F --> G[Email]
    F --> H[Log]
    D --> I[Encrypt Password]
    I --> J[Compare Password]
    J --> K[Success]
    K --> L[Error]
```
This overview shows the relationships between the different parts of the Go file, including the database, creating a new website record, encrypting the website password, comparing the password, and logging any errors. The arrows represent the flow of data or control between these parts, with the source of the arrow indicating the function or method that is being called, and the target indicating the function or method that is being called or the data that is being passed.
Note that this is just a high-level overview, and there may be additional details or nuances in the code that are not captured here.

```
