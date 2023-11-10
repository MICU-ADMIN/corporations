package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "mosque.icu.corporations.api/serv"
)

// GET: Retrieve data from a specified resource.
// Example: Fetching information from a website.

// POST: Submit data to be processed to a specified resource.
// Example: Submitting a form on a website.

// PUT: Update a resource or create a new resource if it doesn't exist.
// Example: Modifying the details of an existing database entry.

// DELETE: Request the removal of a resource.
// Example: Deleting a user account.

// HEAD: Retrieve only the headers of a resource, without the actual data.
// Example: Checking if a resource has been modified.

// OPTIONS: Get information about the communication options available for a resource.
// Example: Checking which HTTP methods are supported.

// PATCH: Apply partial modifications to a resource.
// Example: Making small updates to a document.


func main() {

	router := gin.Default()

	//  open routes	

		router.GET("/register", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// register a user 
		})

		router.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// authenticate a user
		})

	// router.GET("/:name/calendar", func(c *gin.Context) {
	//     id := c.Param("name")
	// 	c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
	// 			//  return ics calendar for this mosque 
	// 		})


	// 		router.GET("/:name/updates", func(c *gin.Context) {
	//     id := c.Param("name")
	// 	c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
	// 			//  return xml rss feed of their updates
	// 		})




// protected by jwt

	// private := router.Group("/private").Use(middleware.CheckJWT())

	// private.POST("/organisation", func(c *gin.Context) {
	// 			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 			//  create a organisation belonging to the yser
	// 		})

	// private.PUT("/api", func(c *gin.Context) {
	// 			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 			//  refresh valid api key for user or create one if it doesnt exist api key for user actions & webhooks
	// 		})

	// private.HEAD("/api", func(c *gin.Context) {
	// 			c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 			//  get the users api key usage could be for user actions or tools such as hadith
	// 		})


	// private.GET("/user", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 		//  retreive data for said user 
	// 	})

	// private.PATCH("/user", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 		//  update user data
	// 	})

	// 	private.DELETE("/user", func(c *gin.Context) {
	// 		c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// 		//  initiate the user deletion process use this as a proof of intention from the user
	// 	})

		// protected by relation within user record and organisation table
	// organisation := router.Group("/organisation").Use(middleware.CheckORG())
// accepts jwt and checks if user is part of org and their role  to see if they have perms 


// 	organisation.PATCH("/", func(c *gin.Context) {
// 				c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
// 				//  update organisation
// 			})

// 	organisation.PUT("/domain", func(c *gin.Context) {
// 					c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
// 					//  refresh current domain for organisation or create one if it doesnt exist 
// 				})

// 	organisation.PUT("/api", func(c *gin.Context) {
// 				c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
// 				//  refresh valid api key for user or create one if it doesnt exist  api key for user actions & webhooks
// 			})

// 	organisation.HEAD("/api", func(c *gin.Context) {
// 				c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
// 				//  get the organisation api key usage
// //  such as 
// 			})

// 				organisation.OPTIONS("/api", func(c *gin.Context) {
// 				c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
// 				//  get all api keys for the org 
// 			})

//    organisation.POST("organisation/:name/updates", func(c *gin.Context) {
// 	    id := c.Param("name")
// 		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
// 				// add a new update from the organisation
// 			})

//    organisation.POST("organisation/:name/event", func(c *gin.Context) {
// 	    id := c.Param("name")
// 		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
// 					// add a new event from the organisation
// 			})

// 	organisation.PATCH("organisation/:name/event", func(c *gin.Context) {
// 	    id := c.Param("name")
// 		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
// 					// update a  event from the organisation
// 			})

// 				organisation.PUT("organisation/:name/prayers", func(c *gin.Context) {
// 	    id := c.Param("name")
// 		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
// 					//  bulk add a prayer time table for a given month or update it if needed on database
// 			})

// 		organisation.POST("organisation/:name/prayer", func(c *gin.Context) {
// // prayer name and time to be updated to in search params 
// 			id := c.Param("name")
// 		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
// 					//   add a prayer time table for a given date 
// 			})


	// protected by relation within user record and organisation table

	establishment := router.Group("/establishment")

// accepts jwt and checks if user is part of org and their role  to see if they have perms 

	establishment.PATCH("organisation/:name/event", func(c *gin.Context) {
	    id := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"data": "Organization ID: " + id}) 
					// update a  event from the organisation
			})


		// protected by api key
	// api := router.Group("/api").Use(middleware.CheckAPI())
	
	router.GET("/hadith", service.GetHadith)

	router.GET("/ahadith", service.GetAhadith)

	router.GET("/ahadith/filter", service.GetAhadithFiltered)



// protected by secret
	// machine := router.Group("/machine").Use(middleware.CheckSCRT())
	
	// machine.DELETE("/user", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})    
	// })
// deletes said user and all associated data 


	// websites
	// Serve static files (HTML, CSS, etc.) for wildcard subdomain
	// router.GET("/:subdomain/*any", func(c *gin.Context) {
	// 	subdomain := c.Param("subdomain")
	// 	filePath := "./web/" + subdomain + "/out/" + c.Param("any")
	// 	c.File(filePath)
	// })


	router.Run()	
}
