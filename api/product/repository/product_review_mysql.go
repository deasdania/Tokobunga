package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type IProductReviewMysql interface {
	GetProductReviewById(id string) (*models.ProductReview, error)
	GetProductReviewByUserId(user_id string) ([]*models.ProductReview, error)
	GetProductReviewByProductId(product_id string) ([]*models.ProductReview, error)
	GetProductReviewByProductIdAndUserId(product_id string, user_id string) (*models.ProductReview, error)
	CreateProductReview(prodrev *models.ProductReview) error
	DeleteProductById(id string) error
}
