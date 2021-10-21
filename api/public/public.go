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

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

// GetProduct godoc
// @Summary Get Product
// @Description Get a list of Product
// @Tags Public
// @Param search query string false "set the product_id or orderby as Query Params"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/product [get]
func (a Public) GetProduct(c *gin.Context) {
	product_id, _ := c.GetQuery("product_id")
	prodid := strings.Trim(product_id, " ")
	orderby, _ := c.GetQuery("orderby")
	orderbytrim := strings.Trim(orderby, " ")

	response := a.ProductUsecase.GetProduct(prodid, orderbytrim)
	c.JSON(response.Status, response)
	return
}

// GetProductReview godoc
// @Summary GetProductReview Public
// @Description free GetProductReview without any auth
// @Tags Public
// @Param Body query string false "set the product_id or user_id or both as Query Params"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/account [get]
func (a Public) GetProductReview(c *gin.Context) {
	user_id, _ := c.GetQuery("user_id")
	product_id, _ := c.GetQuery("product_id")
	response := a.ProductUsecase.GetProductReview(user_id, product_id)
	c.JSON(response.Status, response)
	return
}

// Login godoc
// @Summary Login user
// @Description Logging in to get jwt token to access admin or user api by roles
// @Tags Public
// @Param Body formData LoginInput true "the body to login user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
func (a Public) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	status, token := a.PublicUsecase.Login(email, password)
	c.JSON(status, gin.H{
		"access_token": token,
	})
	return
}

// Register godoc
// @Summary Register Public
// @Description free register without any auth
// @Tags Public
// @Param Body formData models.FormRegister true "fill the form to register"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/create/account [post]
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
