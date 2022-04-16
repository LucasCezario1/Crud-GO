package routes

import (
	"filmes-crud/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") // path ou seja caminho

	movie := main.Group("movies")
	{
		movie.GET("/:id", controllers.ShowMoviebyID)
		movie.GET("/", controllers.ShowMovies)
		movie.POST("/", controllers.CreateMovie)
		movie.PUT("/:id", controllers.UpdateMovie)
		movie.DELETE("/:id", controllers.DeleteMovie)
	}
	return router
}
