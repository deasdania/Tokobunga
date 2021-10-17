package usecase

import (
	accountrepo "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/account/repository"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/product/repository"
	// "strconv"

	// "Final-Project-BDS-Sanbercode-Golang-Batch-28/api/auth/authjwt"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	// "encoding/json"
	// "errors"
	"fmt"
	// "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	"net/url"
	// "os"
	// "strings"
	// "reflect"
)

type productUsecase struct {
	productMysql         repository.IProductMysql
	productCategoryMysql repository.IProductCategoryMysql
	productDetailMysql   repository.IProductDetailMysql
	productSizeMysql     repository.IProductSizeMysql
	responseStruct       response.IResponse
	productReviewMysql   repository.IProductReviewMysql
	accountMysql         accountrepo.IAccountMysql
}

func (a productUsecase) GetProduct(productid string, order string) *response.Response {
	fmt.Print("Get Product", productid)
	if productid == "" {
		allProduct, err := a.productMysql.GetAllProduct(order)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{"cant get all products"}, nil)
		}
		fmt.Println(allProduct)
		fmt.Println(&allProduct)
		return a.responseStruct.ResponseSuccess(200, []string{"All Product"}, map[string][]*models.Product{
			"product": allProduct,
		})
	}
	aProduct, err := a.productMysql.GetProductById(productid)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{"cant get all products"}, nil)
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Get by Product"}, map[string]models.Product{
		"product": *aProduct,
	})
}
func (a productUsecase) CheckReq(productReq *models.ProductReq) *response.Response {
	fmt.Println("productReq", productReq)
	for _, u := range productReq.UrlImage {
		_, err := url.ParseRequestURI(u)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{"url gallery not valid :" + u}, nil)
		}
	}
	// cek apakah produk dengan yang sama ada di db
	_, errs := a.productMysql.GetProductByName(&productReq.Name)
	if errs == nil {
		return a.responseStruct.ResponseError(400, []string{"has been created before, you could do an update"}, nil)
	}
	// cek kategori dengan id yang dikirim apakah ada di db
	_, errGetProdCat := a.productCategoryMysql.GetProductCategoryById(&productReq.CategoryId)
	if errGetProdCat != nil {
		return a.responseStruct.ResponseError(400, []string{"Product Category Error: " + errGetProdCat.Error()}, nil)
	}
	// cek kategori dengan id yang dikirim apakah ada di db
	_, errGetProdSize := a.productSizeMysql.GetProductSizeById(productReq.SizeId)
	if errGetProdSize != nil {
		return a.responseStruct.ResponseError(400, []string{"Product Size Error: " + errGetProdSize.Error()}, nil)
	}
	return a.responseStruct.ResponseError(200, []string{"pass the checking req"}, nil)
}

func (a productUsecase) CreateProduct(productReq *models.ProductReq) *response.Response {
	resChecking := a.CheckReq(productReq)
	if resChecking.Status != 200 {
		return resChecking
	}
	// create the produk
	errCreateProd := a.productMysql.CreateProductReq(productReq)
	if errCreateProd != nil {
		return a.responseStruct.ResponseError(400, []string{errCreateProd.Error()}, nil)
	}

	productcreated, errs := a.productMysql.GetProductByName(&productReq.Name)
	if errs != nil {
		fmt.Println(errs.Error())
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Created Product with details, and galleries"}, map[string]string{
		"product_id":   fmt.Sprintf("%d", productcreated.Id),
		"product_name": productcreated.Name,
	})
}

func (a productUsecase) UpdateProduct(productReq *models.ProductReq) *response.Response {
	resChecking := a.CheckReq(productReq)
	if resChecking.Status != 200 {
		return resChecking
	}
	return a.responseStruct.ResponseError(400, []string{"has been created before, you could do an update"}, nil)
}

func NewProductUsecase(productMysql repository.IProductMysql, productCategoryMysql repository.IProductCategoryMysql, productDetailMysql repository.IProductDetailMysql, productSizeMysql repository.IProductSizeMysql, responseStruct response.IResponse, productReviewMysql repository.IProductReviewMysql, accountMysql accountrepo.IAccountMysql) IProductUsecase {
	return &productUsecase{productMysql: productMysql, productCategoryMysql: productCategoryMysql, productDetailMysql: productDetailMysql, productSizeMysql: productSizeMysql, responseStruct: responseStruct, productReviewMysql: productReviewMysql, accountMysql: accountMysql}
}
