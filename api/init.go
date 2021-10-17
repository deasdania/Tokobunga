package api

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/repository"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/authjwt"
	usecaseauth "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product"
	repoproduct "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/repository"
	usecaseproduct "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role"
	reporole "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/repository"
	usecaserole "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/role/usecase"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/config"
	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/middlewares"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/utilities"
	"fmt"
	"github.com/gin-gonic/gin"
	// "os"
)

func Init(r *gin.Engine) {
	db := config.InitDb()
	redisDb := config.InitDbRedis()

	r.Use(utilities.CORSMiddleware())
	private := r.Group("/api")

	public := r.Group("/api")
	// product := r.Group("/api/product")

	responseStruct := response.InitResponse()
	accountMysql := repository.NewAccountMysql(db)
	roleMysql := reporole.NewroleMysql(db)
	productMysql := repoproduct.NewProductMysql(db)
	productCategoryMysql := repoproduct.NewProductCategoryMysql(db)
	productDetailMysql := repoproduct.NewProductDetailMysql(db)
	productSizeMysql := repoproduct.NewProductSizeMysql(db)

	roleUsecase := usecaserole.NewRoleUsecase(roleMysql, responseStruct)
	accountUsecase := usecase.NewAccountUsecase(roleMysql, accountMysql, responseStruct)
	productUsecase := usecaseproduct.NewProductUsecase(productMysql, productCategoryMysql, productDetailMysql, productSizeMysql, responseStruct)
	//AUTH
	authService := authjwt.JWTAuthService(redisDb)
	authUsecase := usecaseauth.NewAuthUsecase(authService, accountMysql)
	authController := auth.Auth{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	authController.Account(public)
	private.Use(utilities.CheckRestClientJWT(authUsecase))

	//account
	accountController := account.Account{AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	accountController.Account(private)

	//role
	roleController := role.Role{RoleUsecase: roleUsecase, AuthUsecase: authUsecase, AccountUsecase: accountUsecase}
	roleController.Role(private)

	//product
	productController := product.Product{ProductUsecase: productUsecase, AccountUsecase: accountUsecase, AuthUsecase: authUsecase}
	// product.Use(middlewares.JwtAuthMiddleware())
	productController.Product(private)

	fmt.Println(utilities.ACCOUNT_PORT)
	r.Run(fmt.Sprintf(":8089"))
	// r.Run()
}
