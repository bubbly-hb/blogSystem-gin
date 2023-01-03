package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(25);not null" json:"name"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	Email    string `gorm:"type:varchar(100);not null" json:"email"`
}
