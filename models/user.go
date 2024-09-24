package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"uniqueIndex" json:"username"`
	Fullname  string    `json:"fullname"`
    Password string
    Reports []Report `gorm:"foreignKey:AuthorID"`
}