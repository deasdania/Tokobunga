package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type productCategoryMysql struct {
	db *gorm.DB
}

func (a productCategoryMysql) GetProductCategoryById(id *int) (*models.ProductCategory, error) {
	var category models.ProductCategory
	err := a.db.Debug().Model(&models.ProductCategory{}).First(&category, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &category, nil
}

func (a productCategoryMysql) GetProductCategoryByName(name string) (*models.ProductCategory, error) {
	var category models.ProductCategory
	err := a.db.Debug().Model(&models.ProductCategory{}).First(&category, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &category, nil
}

func (a productCategoryMysql) GetAllProductCategory(orderby string) ([]*models.ProductCategory, error) {
	category := make([]*models.ProductCategory, 0)
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
	err := a.db.Debug().Model(&models.ProductCategory{}).Order(orderby).Find(&category)
	if err.Error != nil {
		return nil, err.Error
	}
	return category, nil
}

func (a productCategoryMysql) CreateProductCategory(form *models.FormName) error {
	prodcat := models.ProductCategory{
		Name: form.Name,
	}
	return a.db.Debug().Model(&models.ProductCategory{}).Create(&prodcat).Error
}

func (a productCategoryMysql) UpdateProductName(id string, name string) error {
	return a.db.Debug().Model(&models.ProductCategory{}).Where("id = ?", id).Update("name", name).Error
}

func (a productCategoryMysql) DeleteProductById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.ProductCategory{}).Error
}

func NewProductCategoryMysql(db *gorm.DB) IProductCategoryMysql {
	return &productCategoryMysql{db: db}
}
