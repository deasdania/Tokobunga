package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	"fmt"
	"strconv"
)

func (a productUsecase) CreateProductCategory(form_name models.FormName) *response.Response {
	_, err := a.productCategoryMysql.GetProductCategoryByName(form_name.Name)
	if err == nil {
		return a.responseStruct.ResponseError(400, []string{"has been created before, you could do an update"}, nil)
	}
	err = a.productCategoryMysql.CreateProductCategory(&form_name)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	prodcat, _ := a.productCategoryMysql.GetProductCategoryByName(form_name.Name)
	return a.responseStruct.ResponseError(200, []string{"Create Product Category"}, map[string]string{
		"id":   fmt.Sprintf("%d", prodcat.Id),
		"name": prodcat.Name,
	})
}

func (a productUsecase) UpdateProductCategory(prodcat models.ProductCategory) *response.Response {
	prodcatid := strconv.Itoa(prodcat.Id)
	prodCategory, err := a.productCategoryMysql.GetProductCategoryById(prodcatid)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	if prodcat.Name == prodCategory.Name {
		return a.responseStruct.ResponseError(400, []string{"No changes"}, nil)
	}
	err = a.productCategoryMysql.UpdateProductName(prodcatid, prodcat.Name)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	prodCategory, _ = a.productCategoryMysql.GetProductCategoryById(prodcatid)
	return a.responseStruct.ResponseError(200, []string{"Update Product Category"}, map[string]string{
		"id":   fmt.Sprintf("%d", prodCategory.Id),
		"name": prodCategory.Name,
	})
}
