package port

import (
	"github.com/gin-gonic/gin"
	"pet/internal/handler"
	"pet/internal/service"
)

type Server struct {
	port string
	app  *gin.Engine
}

func NewServer(port string, app *service.Service) Server {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.SignUp(app))
		auth.POST("/sign-in", handler.SignIn(app))
	}

	s := Server{
		port: port,
		app:  router,
	}
	return s
}

func (s *Server) Listen() error {
	return s.app.Run("localhost:" + s.port)
}
