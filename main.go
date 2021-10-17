package main

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/docs"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// for load godotenv
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file not FOUND")
	}
	r := gin.Default()

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Toko Bunga API"
	docs.SwaggerInfo.Description = "This is a Final Project of Camp"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	api.Init(r)
}
