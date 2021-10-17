package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/repository"
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	productrepo "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/repository"
	// rolerepo "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/repository"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	"fmt"
	// "github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	// "regexp"
	// "strconv"
)

type publicUsecase struct {
	accountMysql   repository.IAccountMysql
	responseStruct response.IResponse
	productMysql   productrepo.IProductMysql
}

func (a publicUsecase) CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a publicUsecase) Login(email string, password string) (int, string) {
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
		return 400, err.Error()
	}
	match := a.CheckPasswordHash(password, user.Password)

	if match {
		tokenString, errs := token.GenerateToken(user)
		if errs != nil {
			fmt.Println(err.Error())
			return 400, errs.Error()
		}
		return 200, tokenString
	}
	return 400, "password doesn't match"

}

func NewPublicUsecase(accountMysql repository.IAccountMysql, responseStruct response.IResponse, productMysql productrepo.IProductMysql) IPublicUsecase {
	return &publicUsecase{accountMysql: accountMysql, responseStruct: responseStruct, productMysql: productMysql}
}
