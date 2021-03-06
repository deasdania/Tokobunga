package models

import "time"

// Users
type User struct {
	Id          int       `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	Uuid        string    `json:"uuid" gorm:"unique;not null"`
	Name        string    `json:"name" gorm:"not null;size:150"`
	Email       string    `json:"email" gorm:"unique;not null;size:75"`
	Password    string    `json:"password" gorm:"not null;size:70"`
	CreatedDate time.Time `json:"created_date" gorm:"not null;default:CURRENT_TIMESTAMP;"`
	UpdateDate  time.Time `json:"update_date" gorm:"not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;"`
}

type UserRole struct {
	Id     int `json:"id" gorm:"primary_key;AUTO_INCREMENT;not null`
	UserId int `json:"user_id" gorm:"not null`
	RoleId int `json:"role_id" gorm:"not null`
}

type FormRegister struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type FormChangePassword struct {
	Email           string `json:"email"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}
