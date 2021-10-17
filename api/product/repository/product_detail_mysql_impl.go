package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	// "strings"
)

type productDetailMysql struct {
	db *gorm.DB
}

func (a productDetailMysql) GetProductDetailById(id string) (*models.ProductDetail, error) {
	var detail models.ProductDetail
	err := a.db.Debug().Model(&models.ProductDetail{}).First(&detail, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &detail, nil
}

func (a productDetailMysql) GetProductDetailByProductId(product_id string) (*models.ProductDetail, error) {
	var detail models.ProductDetail
	fmt.Println("ayamZ")
	err := a.db.Debug().Model(&models.ProductDetail{}).First(&detail, "product_id = ?", product_id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &detail, nil
}
func (a productDetailMysql) CreateProductDetail(detail *models.ProductDetail) error {
	return a.db.Debug().Model(&models.ProductDetail{}).Create(&detail).Error
}

// update

func (a productDetailMysql) DeleteProductDetailById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.ProductReview{}).Error
}

func NewProductDetailMysql(db *gorm.DB) IProductDetailMysql {
	return &productDetailMysql{db: db}
}
