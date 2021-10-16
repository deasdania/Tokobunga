package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	"fmt"
	"strconv"
)

func (a productUsecase) CreateProductDetail(detail models.ProductDetail) *response.Response {
	// cek apakah produk dengan yang sama ada di db
	productDetailIdStr := strconv.Itoa(detail.Id)
	detailcreated, errs := a.productDetailMysql.GetProductDetailByProductId(productDetailIdStr)
	if errs == nil {
		return a.responseStruct.ResponseError(400, []string{"has been created before, you could do an update"}, nil)
	}
	// cek product dan size dengan id yang dikirim apakah ada di db
	productIdStr := strconv.Itoa(detail.ProductId)
	_, errGetProd := a.productMysql.GetProductById(productIdStr)
	if errGetProd != nil {
		return a.responseStruct.ResponseError(400, []string{errGetProd.Error()}, nil)
	}
	sizeIdStr := strconv.Itoa(detail.SizeId)
	_, errGetProdSize := a.productSizeMysql.GetProductSizeById(sizeIdStr)
	if errGetProdSize != nil {
		return a.responseStruct.ResponseError(400, []string{errGetProdSize.Error()}, nil)
	}
	// create the produk detail
	errCreateProdDetail := a.productDetailMysql.CreateProductDetail(&detail)
	if errCreateProdDetail != nil {
		return a.responseStruct.ResponseError(400, []string{errCreateProdDetail.Error()}, nil)
	}
	detailcreated, errs = a.productDetailMysql.GetProductDetailByProductId(productDetailIdStr)
	if errs != nil {
		return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Create Product Detail"}, map[string]string{
		"id":         fmt.Sprintf("%d", detailcreated.Id),
		"price":      fmt.Sprintf("%d", detailcreated.Price),
		"product_id": fmt.Sprintf("%d", detailcreated.ProductId),
		"size_id":    fmt.Sprintf("%d", detailcreated.SizeId),
	})
}

//
//func (a productUsecase) UpdateProductDetail(product models.ProductDetail) *response.Response {
//	// cek apakah produk dengan yang sama ada di db
//	intId := strconv.Itoa(product.Id)
//	productg, errs := a.productDetailMysql.GetProductById(intId)
//	if errs != nil {
//		return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
//	}
//	if product.Name == productg.Name {
//		if product.CategoryId == productg.CategoryId {
//			return a.responseStruct.ResponseError(400, []string{"No changes"}, nil)
//		}
//	}
//	catId := strconv.Itoa(product.CategoryId)
//	if product.Name != productg.Name {
//		errs = a.productDetailMysql.UpdateProductName(intId, product.Name)
//		if errs != nil {
//			return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
//		}
//	}
//
//	category, errGetProdCat := a.productCategoryMysql.GetProductCategoryById(catId)
//	if errGetProdCat != nil {
//		return a.responseStruct.ResponseError(400, []string{errGetProdCat.Error()}, nil)
//	}
//	if product.CategoryId != productg.CategoryId {
//		errs = a.productDetailMysql.UpdateProductCategoryId(intId, catId)
//		if errs != nil {
//			return a.responseStruct.ResponseError(400, []string{errs.Error()}, nil)
//		}
//	}
//	return a.responseStruct.ResponseSuccess(200, []string{"Updated Product"}, map[string]string{
//		"id":            fmt.Sprintf("%d", product.Id),
//		"name":          product.Name,
//		"category":      fmt.Sprintf("%d", category.Id),
//		"category_name": category.Name,
//	})
//}