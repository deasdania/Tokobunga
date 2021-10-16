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
	err := db.Debug().AutoMigrate(&models.User{}, &models.UserRole{}, &models.Role{}, &models.Product{}, &models.ProductCategory{}, &models.ProductReview{}, &models.ProductGallery{}, &models.ProductDetail{}, &models.ProductSize{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	} else {
		db.Model(&models.UserRole{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.UserRole{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

		db.Model(&models.Product{}).AddForeignKey("category_id", "product_categories(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.ProductReview{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.ProductReview{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

		db.Model(&models.ProductDetail{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
		db.Model(&models.ProductDetail{}).AddForeignKey("size_id", "product_sizes(id)", "RESTRICT", "RESTRICT")

		db.Model(&models.ProductGallery{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
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

		bytes, _ := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_ACCOUNT_PASSWORD")), 14)

		var user = models.User{
			Id:          1,
			Uuid:        uuid.New().String(),
			Name:        os.Getenv("ADMIN_ACCOUNT_USERNAME"),
			Email:       os.Getenv("ADMIN_ACCOUNT_EMAIL"),
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

		var categories = []models.ProductCategory{
			models.ProductCategory{
				Id:          1,
				Name:        "Bunga Papan",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.ProductCategory{
				Id:          2,
				Name:        "Handbouquet",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.ProductCategory{
				Id:          3,
				Name:        "Standing Flower",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.ProductCategory{
				Id:          4,
				Name:        "Bunga Meja",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
		}
		for i, _ := range categories {
			err = db.Model(&models.ProductCategory{}).Create(&categories[i]).Error
			if err != nil {
				fmt.Printf("[IGNORE THIS]cannot seed categories table: %v \n", err)
			}
		}
		var size = []models.ProductSize{
			models.ProductSize{
				Id:          1,
				Name:        "2M X 125M",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.ProductSize{
				Id:          2,
				Name:        "2M X 150M",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
			models.ProductSize{
				Id:          3,
				Name:        "Standard",
				CreatedDate: time.Now(),
				UpdateDate:  time.Now(),
			},
		}
		for i, _ := range size {
			err = db.Model(&models.ProductSize{}).Create(&size[i]).Error
			if err != nil {
				fmt.Printf("[IGNORE THIS]cannot seed size table: %v \n", err)
			}
		}
	}
}
