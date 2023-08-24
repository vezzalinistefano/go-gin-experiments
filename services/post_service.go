package services

import (
	"github.com/vezzalinistefano/zlogger/models"
	"github.com/vezzalinistefano/zlogger/repositories"
)

type PostService struct {
	postRepository	*repositories.PostRepository
}

func GetPostService() PostService{
	postRepository := repositories.GetInstance()

	return PostService{postRepository: postRepository}
}

func (pr *PostService) Create(newPost models.Post) models.Post{
	return pr.postRepository.Create(newPost)
}

func (pr *PostService) GetAll() []models.Post{
	return pr.postRepository.GetAll()
}