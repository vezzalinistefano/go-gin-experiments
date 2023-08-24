package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vezzalinistefano/zlogger/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post/create", handlers.CreatePost)
	r.GET("/post/getAll", handlers.GetAllPosts)

	r.Run()
}
