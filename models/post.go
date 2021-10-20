package models

type Post struct {
	ID     uint   `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Text   string `json:"text"`
	User   uint   `json:"user_id"`
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
