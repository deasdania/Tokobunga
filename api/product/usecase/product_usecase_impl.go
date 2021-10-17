package usecase

import (
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
}

func (a productUsecase) CreateProduct(productReq *models.ProductReq) *response.Response {
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

// func (a productUsecase) UpdateProduct(product models.Product) *response.Response {
// 	// cek apakah produk dengan yang sama ada di db
// 	intId := strconv.Itoa(product.Id)
// 	productg, errs := a.productMysql.GetProductById(intId)
// 	if errs != nil {
// 		return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
// 	}
// 	if product.Name == productg.Name {
// 		if product.CategoryId == productg.CategoryId {
// 			return a.responseStruct.ResponseError(400, []string{"No changes"}, nil)
// 		}
// 	}
// 	catId := strconv.Itoa(product.CategoryId)
// 	if product.Name != productg.Name {
// 		errs = a.productMysql.UpdateProductName(intId, product.Name)
// 		if errs != nil {
// 			return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
// 		}
// 	}

// 	category, errGetProdCat := a.productCategoryMysql.GetProductCategoryById(catId)
// 	if errGetProdCat != nil {
// 		return a.responseStruct.ResponseError(400, []string{errGetProdCat.Error()}, nil)
// 	}
// 	if product.CategoryId != productg.CategoryId {
// 		errs = a.productMysql.UpdateProductCategoryId(intId, catId)
// 		if errs != nil {
// 			return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
// 		}
// 	}
// 	return a.responseStruct.ResponseSuccess(200, []string{"Updated Product"}, map[string]string{
// 		"id":            fmt.Sprintf("%d", product.Id),
// 		"name":          product.Name,
// 		"category":      fmt.Sprintf("%d", category.Id),
// 		"category_name": category.Name,
// 	})
// }

func NewProductUsecase(productMysql repository.IProductMysql, productCategoryMysql repository.IProductCategoryMysql, productDetailMysql repository.IProductDetailMysql, productSizeMysql repository.IProductSizeMysql, responseStruct response.IResponse) IProductUsecase {
	return &productUsecase{productMysql: productMysql, productCategoryMysql: productCategoryMysql, productDetailMysql: productDetailMysql, productSizeMysql: productSizeMysql, responseStruct: responseStruct}
}
