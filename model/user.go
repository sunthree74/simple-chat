package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string  `json:"name"`
	Phone      string  `json:"phone"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "users"
}
