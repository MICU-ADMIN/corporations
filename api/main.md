```go

package main

import (
	"go-jwt/controllers"
	"go-jwt/database"
	_ "go-jwt/docs"
	"go-jwt/middlewares"
	"go-jwt/models"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// HTTP (Hypertext Transfer Protocol) defines several methods or verbs that indicate the desired action to be performed on a resource. Here are some of the common HTTP methods:
// 1. **GET**: The GET method requests a representation of the specified resource. Requests using GET should only retrieve data and should not have any other effect.
// 2. **POST**: The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server.
// 3. **PUT**: The PUT method replaces all current representations of the target resource with the request payload.
// 4. **DELETE**: The DELETE method deletes the specified resource.
// 5. **PATCH**: The PATCH method applies partial modifications to a resource.
// 6. **HEAD**: The HEAD method asks for the response identical to that of a GET request, but without the response body. This is useful for retrieving meta-information written in response headers, without having to transport the entire content.
// 7. **OPTIONS**: The OPTIONS method is used to describe the communication options for the target resource.
// 8. **TRACE**: The TRACE method performs a message loop-back test along the path to the target resource, providing a useful debugging mechanism.
// 9. **CONNECT**: The CONNECT method establishes a tunnel to the server identified by the target resource.
// These methods provide a wide range of functionalities for interacting with resources on the web. The most commonly used methods are GET and POST, but the others are important in various scenarios, such as updating resources (PUT), deleting resources (DELETE), and applying partial updates (PATCH). The choice of method depends on the desired action and the semantics of the operation you want to perform on the server.

// main is the entry point of the program.
// It initializes the database, sets up the router and starts the server.

// @title Swagger JWT API
// @version 1.0
// @description Create  Go REST API with JWT Authentication in Gin Framework
// @contact.name API Support
// @termsOfService demo.com
// @contact.url http://demo.com/support
// @contact.email support@swagger.io
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /api
//@Schemes http https
// @query.collection.format multi
// @securityDefinitions.basic BasicAuth
// @securityDefinitions.apikey  ApiKeyAuth
// @in header
// @name Authorization
func main() {

   // Initialize the database
 err := database.InitDatabase()

 if err != nil {

   // Log the error and exit
  log.Fatalln("could not create database", err)
 }

 // Automigrate the User model
// AutoMigrate() automatically migrates our schema, to keep our schema upto date.

database.GlobalDB.AutoMigrate(&models.User{})
database.GlobalDB.AutoMigrate(&models.Organization{})
database.GlobalDB.AutoMigrate(&models.Display{})
database.GlobalDB.AutoMigrate(&models.Event{})
database.GlobalDB.AutoMigrate(&models.Member{})
database.GlobalDB.AutoMigrate(&models.Prayer{})
database.GlobalDB.AutoMigrate(&models.Website{})


// Set up the router
 r := setupRouter()

 // Start the server
 r.Run(":8080")
}

// setupRouter sets up the router and adds the routes.
func setupRouter() *gin.Engine {

   // Create a new router
 r := gin.Default()

 // Add a welcome route
 r.GET("/", func(c *gin.Context) {
  c.String(200, "Welcome To This Website")
 })

 // Create a new group for the API
 api := r.Group("/api")
 {

     // Add the login route
   api.POST("/login", controllers.Login)

   // Add the signup route
   api.POST("/signup", controllers.Signup)

   // Create a new group for the apis routes
  // public := api.Group("/public")
  // {

    // get a single hadith
  // public.POST("/hadith", controllers.Login)
 
  // get all hadith
  // public.POST("/ahadith", controllers.Login)
  
  // get hadith by books anc collections
  // public.POST("/ahadith:/filter", controllers.Login)
  // }

//  community := api.Group("/community")
//   {
// //  get community articles 
//   community.GET("/articles", controllers.Login)
 
// // 
//   community.POST("/login", controllers.Login)
 

//   community.POST("/login", controllers.Login)
//   }


// ####################

  // Authentication protected routes
  auth := api.Group("/auth").Use(middlewares.Authz())
  {
   // Get user profile information
   auth.GET("/profile", controllers.Retrieve_Profile)
  
   // Retrieval of the organization
   auth.GET("/organization", controllers.Retrieve_Organization)   

  //  creation of the organization
   auth.POST("/organization", controllers.Create_Organization)
   
  //  // creation of the invitation
  auth.POST("/organization/invite", controllers.Invite_Organization)
 
    //  // the response of the invitation
    auth.PUT("/organization/invite", controllers.Accept_Organization)


  //  protected.POST("/invite/organisation", controllers.Create_Invitation)

  //  protected.PATCH("/invite/organisation", controllers.Respond_Invitation)

  //  // revoking of invitation
  //  protected.PUT("/invite/organisation", controllers.Revoke_Invitation)

  //  // status of invitations
  //  protected.HEAD("/invite/organisation", controllers.Status_Invitation)

   // ###################



   // ###################

  //  // creation of the organisation
  //  auth.POST("/invite/organisation", controllers.Create_Organisation)


// }

  // Add the signup route
  // organisation := api.Group("/oganisation").Use(middlewares.Authz())
  // {
   // Get user profile information
  //  organisation.GET("/profile", controllers.Profile)

  // //  // creation of the organisation
  //  protected.POST("/organisation", controllers.Create_Organisation)


  //  // creation of the invitation
  //  protected.POST("/invite/organisation", controllers.Create_Invitation)

  //  // the response of the email
  //  protected.PATCH("/invite/organisation", controllers.Respond_Invitation)

  //  // revoking of invitation
  //  protected.PUT("/invite/organisation", controllers.Revoke_Invitation)

  //  // status of invitations
  //  protected.HEAD("/invite/organisation", controllers.Status_Invitation)

   // ###################



   // ###################

  //  // creation of the organisation
  //  protected.POST("/invite/organisation", controllers.Create_Organisation)


  //  protected.POST("/organisation", func(c *gin.Context) {
				// c.JSON(http.StatusOK, gin.H{"data": "this is the register endpoint!"})
				//  create a organisation belonging to the yser
			// })
}


}

// docs route
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// Return the router
 return r

}


```

```mermaid

Start the server
r.Run(":8080")
```

This Go file defines a REST API with JWT authentication using the Gin framework. It includes routes for logging in and out, as well as routes for retrieving and creating resources. The `main` function initializes the database, sets up the router, and starts the server. The `setupRouter` function sets up the router and adds the routes, and the `main` function runs the server on port 8080. The `docs` route is used to serve the Swagger documentation for the API.
To generate Mermaid Markdown syntax for this Go file, you can use the following command:
```mermaid
graph LR
    A[main] --> B[database.InitDatabase]
    B --> C[err]
    C --> D[log.Fatalln]
    D --> E[database.GlobalDB.AutoMigrate]
    E --> F[models.User{}]
    F --> G[models.Organization{}]
    G --> H[models.Display{}]
    H --> I[models.Event{}]
    I --> J

```
