package server

import (
	"github.com/gin-gonic/gin"
	"pet/internal/service"
)

type Server struct {
	Hub  *Hub
	port string
	App  *gin.Engine
}

func NewServer(port string, app *service.Service, hub *Hub) Server {
	router := gin.Default()

	//routing(app, router, hub)

	s := Server{
		port: port,
		App:  router,
		Hub:  hub,
	}
	return s
}

func (s *Server) Listen() error {
	return s.App.Run("localhost:" + s.port)
}
