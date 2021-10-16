package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type productMysql struct {
	db *gorm.DB
}

func (a productMysql) GetProductByName(name string) (*models.Product, error) {
	var product models.Product
	err := a.db.Debug().Model(&models.Product{}).First(&product, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &product, nil
}

func (a productMysql) GetProductById(id string) (*models.Product, error) {
	var product models.Product
	err := a.db.Debug().Model(&models.Product{}).First(&product, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &product, nil
}

func (a productMysql) GetAllProduct(orderby string) ([]*models.Product, error) {
	products := make([]*models.Product, 0)
	fmt.Println("orderby")
	fmt.Println(orderby)
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
	err := a.db.Debug().Model(&models.Product{}).Order(orderby).Find(&products)
	if err.Error != nil {
		return nil, err.Error
	}
	return products, nil
}

func (a productMysql) CreateProduct(name *models.FormName) error {
	return a.db.Debug().Model(&models.Product{}).Create(&name).Error
}

func (a productMysql) UpdateProductName(id string, name string) error {
	return a.db.Debug().Model(&models.Product{}).Where("id = ?", id).Update("name", name).Error
}

func (a productMysql) DeleteProductById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.Product{}).Error
}

func NewProductMysql(db *gorm.DB) IProductMysql {
	return &productMysql{db: db}
}
