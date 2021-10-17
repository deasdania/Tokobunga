package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
)

type IRoleUsecase interface {
	GetRoles(id string, orderby string) *response.Response
	CheckRoleExist(id string) bool
	CreateRole(form_name models.FormName) *response.Response
	UpdateRole(role models.Role) *response.Response
}
