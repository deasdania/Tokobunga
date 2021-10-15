package account

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	authusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/usecase"

	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"reflect"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
	AuthUsecase    authusecase.IAuthUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.CREATE_ACCOUNT, a.CreateAccount)
	r.POST(utilities.CHANGE_PASSWORD, a.ChangePassword)

	r.GET(utilities.GENERATE_UUID, a.GenerateUuid)
	r.POST("/test", func(c *gin.Context) { return })
}

func (a Account) GetUser(c *gin.Context) {
	user := a.AccountUsecase.GetUser("1")
	fmt.Println(reflect.TypeOf(user))
	c.JSON(user.Status, user)
}

func (a Account) CreateAccount(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
	}
	fmt.Println(metadata)
	isAdmin := a.AccountUsecase.CheckUserIsAdmin(metadata.Email)
	if isAdmin {
		name := c.PostForm("name")
		email := c.PostForm("email")
		password := c.PostForm("password")
		confirm_password := c.PostForm("confirm_password")
		// role := c.PostForm("roles")

		form_register := models.FormRegister{
			Name:            name,
			Email:           email,
			Password:        password,
			ConfirmPassword: confirm_password,
		}
		response := a.AccountUsecase.CreateUser(form_register, utilities.ADMIN)
		c.JSON(response.Status, response)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"message": "you are not allowed",
	})
}

func (a Account) GenerateUuid(c *gin.Context) {
	myuuid := uuid.New().String()
	c.JSON(http.StatusOK, gin.H{
		"uuid": myuuid,
	})
}

func (a Account) ChangePassword(c *gin.Context) {
	metadata, errA := a.AuthUsecase.ExtractTokenMetadata(c.Request)
	if errA != nil {
		fmt.Println(errA.Error())
	}
	email := metadata.Email
	old_password := c.PostForm("old_password")
	new_password := c.PostForm("new_password")
	confirm_password := c.PostForm("confirm_password")

	form_change_pass := models.FormChangePassword{
		Email:           email,
		OldPassword:     old_password,
		NewPassword:     new_password,
		ConfirmPassword: confirm_password,
	}
	response := a.AccountUsecase.ChangePassword(form_change_pass)

	c.JSON(response.Status, response)
}
