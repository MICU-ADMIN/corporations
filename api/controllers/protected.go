package controllers

import (
	"go-jwt/database"
	"go-jwt/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile is a controller function that retrieves the user profile from the database
// based on the email provided in the authorization middleware.
// It returns a 404 status code if the user is not found,
// and a 500 status code if an error occurs while retrieving the user profile.

// @Summary Get User By Token
// @ID GetUserByToken
// @Produce json
// @Tags User
//Param [param_name] [param_type] [data_type] [required/mandatory] [description]
// @Param Authorization header string true "Authorization header using the Bearer scheme"
// @Success 200 {object} string "Success"
// @Failure 400 {string} string "Error"
// @Router /protected/profile [GET]
func Retrieve_Profile(c *gin.Context) {
 // Initialize a user model
 var user models.User
 // Get the email from the authorization middleware
 email, _ := c.Get("email") 
 // Query the database for the user
 result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)
 // If the user is not found, return a 404 status code
 if result.Error == gorm.ErrRecordNotFound {
  c.JSON(404, gin.H{
   "Error": "User Not Found",
  })
  c.Abort()
  return
 }
 // If an error occurs while retrieving the user profile, return a 500 status code
 if result.Error != nil {
  c.JSON(500, gin.H{
   "Error": "Could Not Get User Profile",
  })
  c.Abort()
  return
 }
 // Set the user's password to an empty string
 user.Password = ""
 // Return the user profile with a 200 status code
 c.JSON(200, user)
}

func Create_Organization(c *gin.Context) {
var organization models.Organization
   // Get ownerID from the context
	ownerID, _ := c.Get("email")

	// Assign ownerID to the Organization struct
	organization.OwnerID = ownerID.(string)

	log.Printf("OwnerID: %s", organization.OwnerID)

	 // Print the organization struct for debugging
    log.Printf("Incoming JSON: %+v", organization)


	// Bind the rest of the JSON data to the struct
	err := c.ShouldBindJSON(&organization)

 if err != nil {
  log.Println(err)
  c.JSON(400, gin.H{
   "Error": "Invalid Inputs ",
  })

  c.Abort()
  return
 }
 err = organization.CreateOrganizationRecord()
 if err != nil {
  log.Println(err)
  c.JSON(500, gin.H{
   "Error": "Error Creating User",
  })
  c.Abort()
  return
 }
}


func Retrieve_Organization(c *gin.Context) {
 // Initialize a user model
 var organization models.Organization
 // Get the email from the authorization middleware
 email, _ := c.Get("email") 
 // Query the database for the organization
 result := database.GlobalDB.Where("owner_id = ?", email.(string)).First(&organization)
 // If the organization is not found, return a 404 status code
 if result.Error == gorm.ErrRecordNotFound {
  c.JSON(404, gin.H{
   "Error": "Organization Not Found",
  })
  c.Abort()
  return
 }
 // If an error occurs while retrieving the organization profile, return a 500 status code
 if result.Error != nil {
  c.JSON(500, gin.H{
   "Error": "Could Not Get Organization Profile",
  })
  c.Abort()
  return
 }
 // Set the organization's password to an empty string
//  organization.Password = ""
 // Return the organization profile with a 200 status code
 c.JSON(200, organization)
}



	// 	// Initialize an organization model
// 	var org models.Organization

// 	// Get the owner's ID from the authorization middleware
//  ownerID, _ := c.Get("email") 

// // Parse JSON from request body and bind it to the organization model
// if err := c.ShouldBindJSON(&org); err != nil {
//     c.JSON(400, gin.H{
//         "error": err.Error(),
//     })
//     c.Abort()
//     return
// }

// 	// Check if required parameters are present
// 	if org.Name == "" || org.Email == "" || org.Phone == "" {
// 		c.JSON(400, gin.H{
// 			"Error": "Missing request parameters. Ensure 'name', 'email', and 'phone' are provided.",
// 		})
// 		c.Abort()
// 		return
// 	}

// 	// Set the owner ID
// 	org.OwnerID = ownerID.(string) // Assuming ownerID is of type int

// 	// Use the CreateOrganizationRecord function to save the organization to the database
// 	if err := org.CreateOrganizationRecord(); err != nil {
// 		c.JSON(500, gin.H{
// 			"Error": "Could not create organization",
// 		})
// 		c.Abort()
// 		return
// 	}

// 	// Return the created organization with a 201 status code
	// c.JSON(201, org)




