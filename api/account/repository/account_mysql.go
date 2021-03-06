package repository

import "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"

type IAccountMysql interface {
	GetAccountById(id string) (*models.User, error)
	GetAccountByEmail(email string) (*models.User, error)
	CreateAccount(user *models.User) error
	UpdateAccountPassword(email string, hash string) error
	CreateUserRole(user *models.User, role *models.Role) error
}
