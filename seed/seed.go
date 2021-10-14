package seed

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"

	"time"
)

func Load(db *gorm.DB) {
	err := db.Debug().AutoMigrate(&models.User{}, &models.UserRole{}, &models.Role{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	} else {
		db.Model(&models.UserRole{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserRole{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
		var roles = []models.Role{
			models.Role{
				Id:          1,
				Name:        "admin",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.Role{
				Id:          2,
				Name:        "member",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
		}

		bytes, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ACCOUNT_SERVICE_PASSWORD")), 14)

		var user = models.User{
			Id:          1,
			Uuid:        uuid.New().String(),
			Name:        os.Getenv("ACCOUNT_SERVICE_NAME"),
			Email:       os.Getenv("ACCOUNT_SERVICE_EMAIL"),
			Password:    string(bytes),
			CreatedDate: time.Now(),
			UpdateDate:  time.Now(),
		}

		var userRole = models.UserRole{
			Id:     1,
			UserId: 1,
			RoleId: 1,
		}
		err := db.Model(&models.User{}).Create(&user).Error
		if err != nil {
			fmt.Printf("[IGNORE THIS]cannot seed User table: %v \n", err)
		}
		for i, _ := range roles {
			err = db.Model(&models.Role{}).Create(&roles[i]).Error
			if err != nil {
				fmt.Printf("[IGNORE THIS]cannot seed Roles table: %v \n", err)
			}
		}
		err = db.Model(&models.UserRole{}).Create(&userRole).Error
		if err != nil {
			fmt.Printf("[IGNORE THIS]cannot seed UserRole table: %v \n", err)
		}
	}
}
