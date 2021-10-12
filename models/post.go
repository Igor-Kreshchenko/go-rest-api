package models

import (
  "github.com/jinzhu/gorm"
)

type Post struct {
  ID     uint   `json:"id" gorm:"primary_key"`
  Title  string `json:"title"`
  Author string `json:"author"`
  Text string `json:"text"`
}

type CreatePostInput struct {
  Title  string `json:"title" binding:"required"`
  Author string `json:"author" binding:"required"`
  Text string `json:"text"`
}

type UpdatePostInput struct {
  Title  string `json:"title"`
  Author string `json:"author"`
  Text string `json:"text"`
}