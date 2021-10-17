package models

import "time"

// Role as Group while on django
type Product struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	CategoryId  int       `json:"category_id" gorm:"not null`
	Name        string    `json:"name" gorm:"not null;size:100"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ProductCategory struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Name        string    `json:"name" gorm:"not null;size:100"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ProductReview struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	ProductId   int       `json:"product_id" gorm:"not null`
	UserId      int       `json:"user_id" gorm:"not null`
	Rating      int       `json:"rating" gorm:"not null`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ProductGallery struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	ProductId   int       `json:"product_id" gorm:"not null`
	UrlImage    string    `json:"url_image" gorm:"not null;size:255"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ProductDetail struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	ProductId   int       `json:"product_id" gorm:"not null`
	SizeId      int       `json:"size_id" gorm:"not null`
	Price       int       `json:"price" gorm:"not null`
	Description string    `json:"description" gorm:"not null;size:255"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type ProductSize struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Name        string    `json:"name" gorm:"not null;size:100"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type FormProductDetail struct {
	ProductId   int    `json:"product_id"`
	SizeId      int    `json:"size_id"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type ProductReq struct {
	CategoryId  int      `json:"category_id"`
	Name        string   `json:"product_name"`
	SizeId      int      `json:"size_id"`
	Price       int      `json:"price"`
	Description string   `json:"description"`
	UrlImage    []string `json:"url_images"`
}

type FormUpdateCategory struct {
	CategoryId string `json:"category_id" binding:"required"`
	ProductId  string `json:"product_id" binding:"required"`
}
