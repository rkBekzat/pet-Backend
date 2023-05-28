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

	api := router.Group("/api", handler.UserIdentity(app))
	{
		pets := api.Group("/pets")
		{
			pets.POST("/", handler.AddPet(app))
			pets.POST("/:id", handler.UpdatePet(app))
			pets.GET("/:id", handler.GetById(app))
			pets.GET("/", handler.GetAll(app))
			pets.DELETE("/:id", handler.DeletePet(app))
		}
		article := api.Group("/article")
		{
			article.POST("/")
			article.POST("/:id")
			article.GET("/")
			article.GET("/:id")

		}
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
