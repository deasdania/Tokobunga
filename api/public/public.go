package public

import (
	accountusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	publicusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/public/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	// "strings"
)

type Public struct {
	PublicUsecase  publicusecase.IPublicUsecase
	AccountUsecase accountusecase.IAccountUsecase
}

func (a Public) Public(r *gin.RouterGroup) {
	r.POST(utilities.LOGIN, a.Login)
	r.POST(utilities.CREATE_ACCOUNT_PUBLIC, a.CreateAccount)

}

// Login godoc
// @Summary login a user.
// @Description user login from public access.
// @Tags Auth
// @Param Body body loginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (a Public) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	status, token := a.PublicUsecase.Login(email, password)
	c.JSON(status, gin.H{
		"access_token": token,
	})
	return
}

func (a Public) CreateAccount(c *gin.Context) {
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
	response := a.AccountUsecase.CreateUser(form_register, utilities.MEMBER)
	c.JSON(response.Status, response)
}
