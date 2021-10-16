package models

import "time"

// Role as Group while on django
type Role struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Name        string    `json:"name" gorm:"not null;size:20"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;"`
}

type FormName struct {
	Name string
}
