package port

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet/internal/handler"
	"pet/internal/service"
)

func NewServer(app service.Service) *http.Server {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", handler.SignUp(app))
		auth.POST("/sign-in", handler.SignIn(app))
	}

	s := &http.Server{Addr: "8080", Handler: router}
	return s
}
