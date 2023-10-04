package handler

import (
	"github.com/gin-gonic/gin"
	"pet/internal/service"
)

func CreateRoom(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//userId := ctx.Param("otherUser")

	}
}

func GetAllRooms(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func JoinRoom(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
