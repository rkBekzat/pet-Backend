package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet/internal/service"
)

func AddAddress(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		_, err = app.Auth.GetName(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

	}
}

func SetAddress(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		_, err = app.Auth.GetName(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}
}
