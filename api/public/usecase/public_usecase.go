package usecase

import (
// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
// "Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
)

type IPublicUsecase interface {
	Login(email string, password string) (int, string)
	CheckPasswordHash(password, hash string) bool
}
