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

func (a productUsecase) UpdateProductCategory(category_id string, name string) *response.Response {
	catId, _ := strconv.Atoi(category_id)
	prodcat, err := a.productCategoryMysql.GetProductCategoryById(&catId)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	if prodcat.Name == name {
		return a.responseStruct.ResponseError(400, []string{"no changes"}, nil)
	}
	err = a.productCategoryMysql.UpdateProductCategory(prodcat.Id, name)
	if err != nil {
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	prodcat, _ = a.productCategoryMysql.GetProductCategoryByName(name)
	return a.responseStruct.ResponseError(200, []string{"Update Product Category"}, map[string]string{
		"id":   fmt.Sprintf("%d", prodcat.Id),
		"name": prodcat.Name,
	})
}
