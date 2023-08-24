package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/zlogger/models"
	"github.com/vezzalinistefano/zlogger/services"
)

func CreatePost(c *gin.Context) {
	var newPost models.Post
	
	postService := services.GetPostService()

	if err := c.ShouldBindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	postService.Create(newPost)
	c.JSON(http.StatusCreated, newPost)
}

func GetAllPosts(c *gin.Context){
	postService := services.GetPostService()
	posts := postService.GetAll()

	c.JSON(http.StatusOK, posts)
}
