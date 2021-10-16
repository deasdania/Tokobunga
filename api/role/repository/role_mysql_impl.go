package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	// "fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type roleMysql struct {
	db *gorm.DB
}

func (a roleMysql) GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Model(&models.Role{}).First(&role, "name = ?", name)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) GetRoleById(id string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Model(&models.Role{}).First(&role, "id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) GetRoleByUserId(id string) (*models.Role, error) {
	var role models.Role
	err := a.db.Debug().Model(&models.Role{}).First(&role, "user_id = ?", id)
	if err.Error != nil {
		return nil, err.Error
	}
	return &role, nil
}

func (a roleMysql) CheckUserIsAdmin(user_id string) (bool, error) {
	var role models.UserRole
	err := a.db.Debug().Model(&models.UserRole{}).First(&role, "user_id = ? AND role_id = 1", user_id)
	if err.Error != nil {
		return false, err.Error
	}
	return true, nil
}

func (a roleMysql) GetAllRole(orderby string) ([]*models.Role, error) {
	roles := make([]*models.Role, 0)
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
	err := a.db.Debug().Model(&models.Role{}).Order(orderby).Find(&roles)
	if err.Error != nil {
		return nil, err.Error
	}
	return roles, nil
}

func (a roleMysql) CreateRole(name *models.FormName) error {
	return a.db.Debug().Model(&models.Role{}).Create(&name).Error
}

func (a roleMysql) UpdateRoleName(id string, name string) error {
	return a.db.Debug().Model(&models.Role{}).Where("id = ?", id).Update("name", name).Error
}

func (a roleMysql) DeleteRoleById(id string) error {
	return a.db.Where("id = ?", id).Delete(&models.Role{}).Error
}

func NewroleMysql(db *gorm.DB) IRoleMysql {
	return &roleMysql{db: db}
}
