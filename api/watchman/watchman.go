package watchman

import (
	_ "go-jwt/docs"
	"log"
)

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
func OnUserCreation() {
log.Println("hello")
}