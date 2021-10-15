package models

import "github.com/jinzhu/gorm"

type Post struct {
	gorm.Model
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
	User   User   `gorm:"foreignKey:UserRefer"`
}

type CreatePostInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Text   string `json:"text"`
}

type UpdatePostInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
}
