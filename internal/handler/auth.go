package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"pet/internal/entity"
	"pet/internal/service"
)

func SignUp(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input entity.User
		if err := ctx.BindJSON(&input); err != nil {
			fmt.Println(err.Error())
			log.Println(err.Error())
			return
		}
		fmt.Println(input)
		id, err := app.Auth.CreateUser(input)
		if err != nil {
			fmt.Println(err.Error())
			log.Println(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
	}
}

func SignIn(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input entity.User
		if err := ctx.BindJSON(&input); err != nil {
			return
		}
		fmt.Println(input)
		token, err := app.Auth.GenerateToken(input.Username, input.Password)
		if err != nil {
			fmt.Println(err.Error())
			log.Println(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{"token": token})
	}
}
