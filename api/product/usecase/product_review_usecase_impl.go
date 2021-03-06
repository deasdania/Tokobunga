package usecase

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/response"
	// "errors"
	"fmt"
	"strconv"
)

func (a productUsecase) GetProductReview(user_id string, product_id string) *response.Response {
	fmt.Println("user_id, product_id")
	fmt.Println(user_id, product_id)
	// jika id kosong dan product_id isi
	if user_id != "" && product_id == "" {
		productRev, err := a.productReviewMysql.GetProductReviewByUserId(user_id)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseSuccess(200, []string{"Get Product Review by Id"}, map[string][]*models.ProductReview{
			"product_review": productRev,
		})
	}
	// jika user_id isi dan product_id kosong
	if user_id == "" && product_id != "" {
		productRevs, err := a.productReviewMysql.GetProductReviewByProductId(product_id)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseSuccess(200, []string{"Get Product Review by Product Id"}, map[string][]*models.ProductReview{
			"product_reviews": productRevs,
		})
	}
	// jika user_id isi dan product_id isi
	if user_id != "" && product_id != "" {
		productRev, err := a.productReviewMysql.GetProductReviewByProductIdAndUserId(product_id, user_id)
		if err != nil {
			return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
		}
		return a.responseStruct.ResponseSuccess(200, []string{"Get Product Review by Product Id"}, map[string]*models.ProductReview{
			"product_reviews": productRev,
		})
	}
	return a.responseStruct.ResponseError(400, []string{"please set product_id or user_id or both"}, nil)
}

func (a productUsecase) CreateProductReview(prodRev models.ProductReview, email string) *response.Response {
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	product_idstr := strconv.Itoa(prodRev.ProductId)
	user_idstr := strconv.Itoa(user.Id)

	productRev, errsql := a.productReviewMysql.GetProductReviewByProductIdAndUserId(product_idstr, user_idstr)
	if errsql == nil {
		return a.responseStruct.ResponseError(400, []string{"Product Review by User for the Product Id has been created before, do an update instead"}, map[string]*models.ProductReview{
			"product_reviews": productRev,
		})
	}

	if prodRev.Rating > 5 || prodRev.Rating < 1 {
		return a.responseStruct.ResponseError(400, []string{"Rating couldn't be more than 5 and less than 1"}, nil)
	}
	prodRev.UserId = user.Id
	errCreateProd := a.productReviewMysql.CreateProductReview(&prodRev)
	if errCreateProd != nil {
		return a.responseStruct.ResponseError(400, []string{errCreateProd.Error()}, nil)
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Create Product Review"}, map[string]models.ProductReview{
		"product_reviews": prodRev,
	})
}

func (a productUsecase) UpdateProductReview(prodRev models.ProductReview, email string) *response.Response {
	user, err := a.accountMysql.GetAccountByEmail(email)
	if err != nil {
		fmt.Println(err.Error())
		return a.responseStruct.ResponseError(400, []string{err.Error()}, nil)
	}
	product_idstr := strconv.Itoa(prodRev.ProductId)
	user_idstr := strconv.Itoa(user.Id)

	productRev, errsql := a.productReviewMysql.GetProductReviewByProductIdAndUserId(product_idstr, user_idstr)
	if errsql != nil {
		return a.responseStruct.ResponseError(400, []string{"Product Review by User for the Product Id not in db"}, nil)
	}
	if prodRev.Rating == productRev.Rating {
		return a.responseStruct.ResponseError(400, []string{"No changes"}, nil)
	}
	if prodRev.Rating > 5 || prodRev.Rating < 1 {
		return a.responseStruct.ResponseError(400, []string{"Rating couldn't be more than 5 and less than 1"}, nil)
	}
	prodRev.UserId = user.Id
	productRevIdstr := strconv.Itoa(productRev.Id)

	errUpdateProdRev := a.productReviewMysql.UpdateProductReview(productRevIdstr, prodRev.Rating)
	if errUpdateProdRev != nil {
		return a.responseStruct.ResponseError(400, []string{errUpdateProdRev.Error()}, nil)
	}
	return a.responseStruct.ResponseSuccess(200, []string{"Create Product Review"}, map[string]models.ProductReview{
		"product_reviews": prodRev,
	})
}
