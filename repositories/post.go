package repositories

import (
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	CreatePost(post *models.Post) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	UpdatePostText(id uint, newText string) error
	DeletePost(id uint) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	res := r.db.Find(&posts)

	return posts, res.Error
}

func (r *postRepository) CreatePost(post *models.Post) (*models.Post, error) {
	res := r.db.Create(&post)

	return post, res.Error
}

func (r *postRepository) GetPostByID(id uint) (*models.Post, error) {
	var post *models.Post
	res := r.db.First(&post, id)

	return post, res.Error
}

func (r *postRepository) UpdatePostText(id uint, newText string) error {
	var post *models.Post

	res := r.db.Model(&post).Where("id = ?", id).Update("text", newText)

	return res.Error
}

func (r *postRepository) DeletePost(id uint) error {
	res := r.db.Where(id).Delete(&models.Post{})

	return res.Error
}
