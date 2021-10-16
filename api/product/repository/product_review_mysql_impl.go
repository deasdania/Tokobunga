package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type productReviewMysql struct {
	db *gorm.DB
}

func (a productReviewMysql) GetProductReviewById(id string) (*models.ProductReview, error) {
	var review models.ProductReview
	err := a.db.Debug().Model(&models.ProductReview{}).First(&review, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &review, nil
}

func (a productReviewMysql) GetProductReviewByProductId(product_id string) (*[]models.ProductReview, error) {
	var review models.ProductReview
	err := a.db.Debug().Model(&models.ProductReview{}).Find(&review, "product_id = ?", product_id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &review, nil
}

// create
// update

func (a productReviewMysql) DeleteProductById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.ProductReview{}).Error
}

func NewProductReviewMysql(db *gorm.DB) IProductReviewMysql {
	return &productReviewMysql{db: db}
}
