package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pet/internal/entity"
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
		var address entity.Address
		err = ctx.BindJSON(&address)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		err = app.Owner.AddAddress(address)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok")
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
		var address entity.Address
		err = ctx.BindJSON(&address)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		err = app.Owner.AddAddress(address)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok")
	}
}
