package server

import (
	"github.com/gin-gonic/gin"
	"pet/internal/handler"
	"pet/internal/service"
)

func routing(app *service.Service, router *gin.Engine) {
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
			article.POST("/", handler.CreatePost(app))
			article.POST("/:id", handler.UpdatePost(app))
			article.GET("/", handler.GetAllPost(app))
			article.GET("/:id", handler.GetPost(app))
			article.DELETE("/:id", handler.DeletePost(app))
		}

		api.POST("/", handler.AddAddress(app))
		api.POST("/update", handler.SetAddress(app))

	}
	api.GET("/socket")
}
