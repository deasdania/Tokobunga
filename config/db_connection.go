package config

import (
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/seed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
	"os"
)

var db *gorm.DB //database

func InitDb() *gorm.DB {

	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		username := os.Getenv("DATABASE_USER")
		password := os.Getenv("DATABASE_PASSWORD")
		host := os.Getenv("DATABASE_HOST")
		port := os.Getenv("DATABASE_PORT")
		database := os.Getenv("DATABASE_NAME")
		// production

		dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=require password=%s", host, port, username, database, password) //Build connection string
		fmt.Println(dbUri)

		conn, err := gorm.Open("postgres", dbUri)
		if err != nil {
			fmt.Print(err)
		}

		db = conn
		seed.Load(db)
		return db
	} else {
		dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("DATABASE_USER"),
			os.Getenv("DATABASE_PASSWORD"),
			os.Getenv("DATABASE_HOST"),
			os.Getenv("DATABASE_PORT"),
			os.Getenv("DATABASE_NAME"),
		)
		db, err := gorm.Open("mysql", dataSource)
		if err != nil {
			panic(err)
			return nil
		}

		seed.Load(db)
		return db
	}
}
