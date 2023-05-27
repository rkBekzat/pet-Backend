package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet/internal/entity"
	"pet/internal/service"
)

func SignUp(app service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input entity.User
		if err := ctx.BindJSON(&input); err != nil {
			return
		}
		id, err := app.Auth.CreateUser(input)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
	}
}

func SignIn(app service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input entity.User
		if err := ctx.BindJSON(&input); err != nil {
			return
		}
		token, err := app.Auth.GenerateToken(input)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})
	}
}
