package main

import (
	"fmt"
	"net/http"

	"github.com/EksplorasiGin/controllers"
	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func main() {
	router := gin.Default()

	router.GET("/ping", test)

	router.GET("/movies", controllers.GetAllMovies)
	router.GET("/movie", controllers.GetMovie)
	router.POST("/movie", controllers.InsertMovie)
	router.PUT("/movie", controllers.UpdateMovie)
	router.DELETE("/movie", controllers.DeleteMovie)

	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	router.Run(":8080")
	if err := router.Run(":8080"); err != nil {
		fmt.Println("Router couldnt connect")
	}

}
