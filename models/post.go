package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string `json:"title"`
	Author    string `json:"author"`
	Text      string `json:"text"`
	UserRefer uint   `json:"user_refer"`
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
