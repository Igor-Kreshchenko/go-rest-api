package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
	UserID uint
}
