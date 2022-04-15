package server

import (
	"filmes-crud/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine //server para iniciar o server

}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	routes := routes.ConfigRoutes(s.server)

	log.Print("server esta funcionado " + s.port)
	log.Fatal(routes.Run(":" + s.port))
}
