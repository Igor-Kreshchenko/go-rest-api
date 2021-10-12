package models

type Post struct {
	ID    uint   `json:"id" gorm:"primary_key"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Owner string `json:"owner"`
}