package handler

import (
	"github.com/gin-gonic/gin"
	"pet/internal/server"
	"pet/internal/service"
)

func Routing(app *service.Service, router *gin.Engine, hub *server.Hub) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", SignUp(app))
		auth.POST("/sign-in", SignIn(app))
	}

	api := router.Group("/api", UserIdentity(app))
	{
		pets := api.Group("/pets")
		{
			pets.POST("/", AddPet(app))
			pets.POST("/:id", UpdatePet(app))
			pets.GET("/:id", GetById(app))
			pets.GET("/", GetAll(app))
			pets.DELETE("/:id", DeletePet(app))
		}
		article := api.Group("/article")
		{
			article.POST("/", CreatePost(app))
			article.POST("/:id", UpdatePost(app))
			article.GET("/", GetAllPost(app))
			article.GET("/:id", GetPost(app))
			article.DELETE("/:id", DeletePost(app))
		}

		api.POST("/", AddAddress(app))
		api.POST("/update", SetAddress(app))

		api.POST("/createRoom/:otherUser", CreateRoom(app, hub))
		api.GET("/getRooms", GetAllRooms(app, hub))
		api.GET("/joinRoom/:room_id", JoinRoom(app, hub))
	}
	api.GET("/socket")
}
