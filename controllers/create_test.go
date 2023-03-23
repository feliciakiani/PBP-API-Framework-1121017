package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetAllMovies(t *testing.T) {
	// Initialize a new Gin router
	router := gin.Default()

	// Define the route to test
	router.GET("/movies", GetAllMovies)

	// Create a new HTTP request for the defined route
	req, err := http.NewRequest("GET", "/movies", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a new HTTP response recorder to record the response from the handler
	res := httptest.NewRecorder()

	// Serve the request to the handler
	router.ServeHTTP(res, req)

	// Check the status code of the response
	if res.Code != http.StatusOK {
		t.Errorf("Unexpected status code: got %v want %v", res.Code, http.StatusOK)
	}
}
