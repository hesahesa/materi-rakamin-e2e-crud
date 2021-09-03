package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title string `json:"title"`
}

func MigrateBook(db *gorm.DB) {
	db.AutoMigrate(&Book{})
}