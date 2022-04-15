package routes

import (
	"filmes-crud/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") // path ou seja caminho

	movie := main.Group("movies")
	{
		movie.GET("/id", controllers.ShowMovie)

	}
	return router
}
