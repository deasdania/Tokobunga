package api

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/repository"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/public"
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth"
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/authjwt"
	// usecaseauth "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product"
	repoproduct "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/repository"
	usecaseproduct "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/usecase"
	usecasepublic "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/public/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role"
	reporole "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/repository"
	usecaserole "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/config"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/middlewares"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func Init(r *gin.Engine) {
	db := config.InitDb()

	r.Use(utilities.CORSMiddleware())
	private := r.Group("/api")
	private.Use(middlewares.JwtAuthMiddleware())

	public_api := r.Group("/api")
	// product := r.Group("/api/product")

	responseStruct := response.InitResponse()
	accountMysql := repository.NewAccountMysql(db)
	roleMysql := reporole.NewroleMysql(db)
	productMysql := repoproduct.NewProductMysql(db)
	productCategoryMysql := repoproduct.NewProductCategoryMysql(db)
	productDetailMysql := repoproduct.NewProductDetailMysql(db)
	productSizeMysql := repoproduct.NewProductSizeMysql(db)

	accountUsecase := usecase.NewAccountUsecase(roleMysql, accountMysql, responseStruct)
	roleUsecase := usecaserole.NewRoleUsecase(roleMysql, responseStruct)
	productUsecase := usecaseproduct.NewProductUsecase(productMysql, productCategoryMysql, productDetailMysql, productSizeMysql, responseStruct)
	publicUsecase := usecasepublic.NewPublicUsecase(accountMysql, responseStruct, productMysql)

	//account
	accountController := account.Account{AccountUsecase: accountUsecase}
	accountController.Account(private)

	//role
	roleController := role.Role{RoleUsecase: roleUsecase, AccountUsecase: accountUsecase}
	roleController.Role(private)

	//product
	productController := product.Product{ProductUsecase: productUsecase, AccountUsecase: accountUsecase}
	productController.Product(private)

	//public
	publicController := public.Public{PublicUsecase: publicUsecase, AccountUsecase: accountUsecase}
	publicController.Public(public_api)

	fmt.Println(utilities.ACCOUNT_PORT)
	// r.Run(fmt.Sprintf(":8089"))
	r.Run(os.Getenv("HOST"))
}
