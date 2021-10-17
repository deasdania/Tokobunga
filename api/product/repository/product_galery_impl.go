package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type productGalleryMysql struct {
	db *gorm.DB
}

func (a productGalleryMysql) GetProductGalleryById(id *int) (*models.ProductGallery, error) {
	var category models.ProductGallery
	err := a.db.Debug().Model(&models.ProductGallery{}).First(&category, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &category, nil
}

func (a productGalleryMysql) GetProductGalleryByName(name string) (*models.ProductGallery, error) {
	var category models.ProductGallery
	err := a.db.Debug().Model(&models.ProductGallery{}).First(&category, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &category, nil
}

func (a productGalleryMysql) GetAllProductGallery(orderby string) ([]*models.ProductGallery, error) {
	category := make([]*models.ProductGallery, 0)
	if orderby != "" {
		sortBy := strings.ToUpper(orderby)
		if sortBy == "DESC" || sortBy == "ASC" {
			orderby = "created_date " + orderby
		} else {
			orderby = "created_date desc"
		}
	} else {
		orderby = "created_date desc"
	}
	err := a.db.Debug().Model(&models.ProductGallery{}).Order(orderby).Find(&category)
	if err.Error != nil {
		return nil, err.Error
	}
	return category, nil
}
func (a productGalleryMysql) GetProductGalleryByProductId(product_id string) (*models.ProductGallery, error) {
	var gallery models.ProductGallery
	fmt.Println("ayamZ")
	err := a.db.Debug().Model(&models.ProductGallery{}).First(&gallery, "product_id = ?", product_id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &gallery, nil
}

func (a productGalleryMysql) CreateProductGallery(gallery *models.ProductGallery) error {
	return a.db.Debug().Model(&models.ProductGallery{}).Create(&gallery).Error
}

func (a productGalleryMysql) UpdateProductName(id string, name string) error {
	return a.db.Debug().Model(&models.ProductGallery{}).Where("id = ?", id).Update("name", name).Error
}

func (a productGalleryMysql) DeleteProductGalleryById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.ProductGallery{}).Error
}

func NewProductGalleryMysql(db *gorm.DB) IProductGalleryMysql {
	return &productGalleryMysql{db: db}
}
