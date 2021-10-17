package public

import (
	accountusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	productusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/usecase"
	publicusecase "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/public/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"

	// "fmt"
	"github.com/gin-gonic/gin"
	// "net/http"
	"strings"
)

type Public struct {
	PublicUsecase  publicusecase.IPublicUsecase
	AccountUsecase accountusecase.IAccountUsecase
	ProductUsecase productusecase.IProductUsecase
}

func (a Public) Public(r *gin.RouterGroup) {
	r.POST(utilities.LOGIN, a.Login)
	r.POST(utilities.CREATE_ACCOUNT_PUBLIC, a.CreateAccount)

	r.GET(utilities.GET_PRODUCT, a.GetProduct)
	r.GET(utilities.GET_PRODUCT_REVIEW, a.GetProductReview)
}

func (a Public) GetProduct(c *gin.Context) {
	product_id, _ := c.GetQuery("product_id")
	prodid := strings.Trim(product_id, " ")
	orderby, _ := c.GetQuery("orderby")
	orderbytrim := strings.Trim(orderby, " ")

	response := a.ProductUsecase.GetProduct(prodid, orderbytrim)
	c.JSON(response.Status, response)
	return
}

func (a Public) GetProductReview(c *gin.Context) {
	user_id, _ := c.GetQuery("user_id")
	product_id, _ := c.GetQuery("product_id")
	response := a.ProductUsecase.GetProductReview(user_id, product_id)
	c.JSON(response.Status, response)
	return
}

// Login godoc
// @Summary Login
// @Description Login User
// @Tags Auth
// @Produce json
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
