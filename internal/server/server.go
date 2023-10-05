package server

import (
	"github.com/gin-gonic/gin"
	"pet/internal/service"
)

type Server struct {
	hub  *Hub
	port string
	app  *gin.Engine
}

func NewServer(port string, app *service.Service, hub *Hub) Server {
	router := gin.Default()

	routing(app, router, hub)

	s := Server{
		port: port,
		app:  router,
		hub:  hub,
	}
	return s
}

func (s *Server) Listen() error {
	return s.app.Run("localhost:" + s.port)
}
