package main

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
	// if len(os.Args) > 1 {
	// if len(os.Args) > 1 && os.Args[1] == "local" {
	// }
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file not FOUND")
	}
	r := gin.Default()

	api.Init(r)
}
