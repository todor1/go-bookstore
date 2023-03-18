package models

import (
	"gihub.com/jinzhu/gorm"
	"github.com/jinzhu/gorm"
	"github.com/todor1/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name       string `gorm:""json:"name"`
	Author     string `json: "author"`
	Publicatin string `json: "publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
