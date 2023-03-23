package main

import (
	"github.com/EksplorasiGin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/movies", controllers.GetAllMovies)
	router.GET("/movie", controllers.GetMovie)
	router.POST("/movie", controllers.InsertMovie)
	router.PUT("/movie", controllers.UpdateMovie)
	router.DELETE("/movie", controllers.DeleteMovie)

	router.POST("/login", controllers.Login)
	router.POST("/logout", controllers.Logout)

	_ = router.Run(":8080")

}
