package services

import (
	"github.com/Igor-Kreshchenko/go-rest-api/models"
	"github.com/Igor-Kreshchenko/go-rest-api/repositories"
)

type PostRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Text   string `json:"text" binding:"required"`
	UserID uint   `json:"user_id" binding:"required"`
}

type Text struct {
	Text string `json:"text" binding:"required"`
}

type PostService interface {
	GetAllPosts() ([]models.Post, error)
	CreatePost(post *PostRequest) (*models.Post, error)
	GetPostByID(id uint) (*models.Post, error)
	UpdatePostText(id uint, newText string) error
	DeletePost(id uint) error
}

type postService struct {
	postRepository repositories.PostRepository
}

func NewPostService(postRepo repositories.PostRepository) PostService {
	return &postService{postRepository: postRepo}
}

func (s *postService) GetAllPosts() ([]models.Post, error) {
	res, err := s.postRepository.GetAllPosts()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *postService) CreatePost(postReq *PostRequest) (*models.Post, error) {
	post := models.Post{
		Title:  postReq.Title,
		Author: postReq.Author,
		Text:   postReq.Text,
		UserID: postReq.UserID,
	}

	res, err := s.postRepository.CreatePost(&post)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (s *postService) GetPostByID(id uint) (*models.Post, error) {
	post, err := s.postRepository.GetPostByID(id)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (s *postService) UpdatePostText(id uint, text string) error {
	var newText string

	_, err := s.postRepository.GetPostByID(id)
	if err != nil {
		return err
	}

	newText = text

	err = s.postRepository.UpdatePostText(id, newText)
	if err != nil {
		return err
	}

	return nil
}

func (s *postService) DeletePost(id uint) error {
	err := s.postRepository.DeletePost(id)
	if err != nil {
		return err
	}

	return nil
}
