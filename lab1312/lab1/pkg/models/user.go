package models

import (
	"lab1312lab1/pkg/config"

	"gorm.io/gorm"
)

type User struct {
	Id      uint64 `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func AddUser(us *User) *User {
	db.Create(&us)
	return us
}
