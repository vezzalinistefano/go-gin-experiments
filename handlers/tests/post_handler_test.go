package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/vezzalinistefano/zlogger/handlers"
	"github.com/vezzalinistefano/zlogger/models"
)

func TestCreatePost(t *testing.T) {
	// Set up the Gin router
	router := gin.Default()
	router.POST("/posts", handlers.CreatePost)

	// Create a test request with a JSON payload
	newPost := models.Post{ID: 1, Title: "Test Post", Body: "This is a test post"}
	payload, _ := json.Marshal(newPost)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(payload))

	// Create a response recorder to capture the response
	res := httptest.NewRecorder()

	// Serve the request to the router
	router.ServeHTTP(res, req)

	// Assert that the response status code is 201 Created
	if res.Code != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d.", http.StatusCreated, res.Code)
	}

	// Assert that the response body contains the created post
	var createdPost models.Post
	err := json.Unmarshal(res.Body.Bytes(), &createdPost)

	if err != nil {
		t.Errorf("Can't unmarshal the body response.")
	}

	if !createdPost.Equals(newPost) {
		t.Error("The created post differs from what was intended to be created")
	}
}
