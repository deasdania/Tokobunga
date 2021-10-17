package account

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities/token"
	// "fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	// "reflect"
)

type Account struct {
	AccountUsecase usecase.IAccountUsecase
}

func (a Account) Account(r *gin.RouterGroup) {
	r.GET(utilities.CHECK_AUTH, a.CheckAuth)

	// r.GET(utilities.GET_ACCOUNT, a.GetUser)
	r.POST(utilities.CREATE_ACCOUNT, a.CreateAccount)
	r.POST(utilities.CHANGE_PASSWORD, a.ChangePassword)

	r.GET(utilities.GENERATE_UUID, a.GenerateUuid)
	r.POST("/test", func(c *gin.Context) { return })
}

// used to be use while admin has to create user on the background
func (a Account) CreateAccount(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	confirm_password := c.PostForm("confirm_password")

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

func (a Account) GenerateUuid(c *gin.Context) {
	myuuid := uuid.New().String()
	c.JSON(http.StatusOK, gin.H{
		"uuid": myuuid,
	})
}

// CheckAuth godoc
// @Summary CheckAuth Private
// @Description check token header
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/check/authorize [post]
func (a Account) CheckAuth(c *gin.Context) {
	err := token.TokenValid(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Valid Token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Valid Token",
	})
}

// ChangePassword godoc
// @Summary ChangePassword Private
// @Description change password by user who logged in
// @Tags Private
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body formData models.FormChangePassword true "set form to change your password"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/change/account/password [post]
func (a Account) ChangePassword(c *gin.Context) {
	tokenstring := token.ExtractToken(c)
	metadata, _ := token.ExtractTokenMetadata(tokenstring)
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
