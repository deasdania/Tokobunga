package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
)

type IAccountUsecase interface {
	GetUser(id string) *response.Response
	CheckUserExist(email string) bool
	CheckPasswordLever(ps string) []string
	CreateUser(form_register models.FormRegister) *response.Response
	ChangePassword(form_change_pass models.FormChangePassword) *response.Response
}
