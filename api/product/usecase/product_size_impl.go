package usecase

// import (
// 	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
// 	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
// 	"fmt"
// 	"strconv"
// )

// func (a productUsecase) CreateProductSize(form_name models.FormName) *response.Response {
// 	_, err := a.productSizeMysql.GetProductSizeByName(form_name.Name)
// 	if err == nil {
// 		return a.responseStruct.ResponseError(400, []string{"has been created before, you could do an update"}, nil)
// 	}
// 	err = a.productSizeMysql.CreateProductSize(&form_name)
// 	if err != nil {
// 		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
// 	}
// 	prodSize, _ := a.productSizeMysql.GetProductSizeByName(form_name.Name)
// 	return a.responseStruct.ResponseError(200, []string{"Create Product Size"}, map[string]string{
// 		"id":   fmt.Sprintf("%d", prodSize.Id),
// 		"name": prodSize.Name,
// 	})
// }

// func (a productUsecase) UpdateProductSize(prodsize models.ProductSize) *response.Response {
// 	prodSizeId := strconv.Itoa(prodsize.Id)
// 	prodSize, err := a.productSizeMysql.GetProductSizeById(&prodsize.Id)
// 	if err != nil {
// 		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
// 	}
// 	if prodsize.Name == prodSize.Name {
// 		return a.responseStruct.ResponseError(400, []string{"No changes"}, nil)
// 	}
// 	err = a.productSizeMysql.UpdateProductSizeName(prodSizeId, prodsize.Name)
// 	if err != nil {
// 		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
// 	}
// 	prodSize, _ = a.productSizeMysql.GetProductSizeById(&prodsize.Id)
// 	return a.responseStruct.ResponseError(200, []string{"Update Product Size"}, map[string]string{
// 		"id":   fmt.Sprintf("%d", prodSize.Id),
// 		"name": prodSize.Name,
// 	})
// }
