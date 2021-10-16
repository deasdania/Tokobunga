package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type productSizeMysql struct {
	db *gorm.DB
}

func (a productSizeMysql) GetProductSizeByName(name string) (*models.ProductSize, error) {
	var size models.ProductSize
	err := a.db.Debug().Model(&models.ProductSize{}).First(&size, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &size, nil
}

func (a productSizeMysql) GetProductSizeById(id string) (*models.ProductSize, error) {
	var size models.ProductSize
	err := a.db.Debug().Model(&models.ProductSize{}).First(&size, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &size, nil
}

func (a productSizeMysql) GetAllProductSize(orderby string) ([]*models.ProductSize, error) {
	size := make([]*models.ProductSize, 0)
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
	err := a.db.Debug().Model(&models.ProductSize{}).Order(orderby).Find(&size)
	if err.Error != nil {
		return nil, err.Error
	}
	return size, nil
}

func (a productSizeMysql) CreateProductSize(name *models.FormName) error {
	return a.db.Debug().Model(&models.ProductSize{}).Create(&name).Error
}

func (a productSizeMysql) UpdateProductSizeName(id string, name string) error {
	return a.db.Debug().Model(&models.ProductSize{}).Where("id = ?", id).Update("name", name).Error
}

func (a productSizeMysql) DeleteProductSizeById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.ProductSize{}).Error
}

func NewProductSizeMysql(db *gorm.DB) IProductSizeMysql {
	return &productSizeMysql{db: db}
}
