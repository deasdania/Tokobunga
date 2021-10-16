package repository

import (
	"Final-Project-BDS-Sanbercode-Golang-Batch-28/api/models"
)

type IProductReviewMysql interface {
	GetProductReviewById(id string) (*models.ProductReview, error)
	DeleteProductById(id string) error
}
