package server

import (
	"github.com/gin-gonic/gin"
	"pet/internal/service"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewServer(port string, app *service.Service) Server {
	router := gin.Default()

	routing(app, router)

	s := Server{
		port: port,
		app:  router,
	}
	return s
}

func (s *Server) Listen() error {
	return s.app.Run("localhost:" + s.port)
}
