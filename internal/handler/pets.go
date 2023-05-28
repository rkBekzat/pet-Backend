package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"pet/internal/entity"
	"pet/internal/service"
	"strconv"
)

func AddPet(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		username, err := app.Auth.GetName(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		var pet entity.Pet
		err = ctx.BindJSON(&pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		pet.Username = &username
		fmt.Println(pet.Username, pet.Sex, pet.IssuedOrganization, pet.IssuedOrganization)
		err = app.Pet.Create(pet)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok!")
	}
}

func GetAll(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		username, err := app.Auth.GetName(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		all, err := app.Pet.GetAll(username)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, struct {
			Data []*entity.Pet `json:"data"`
		}{
			Data: all,
		})
	}
}

func GetById(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := ctx.Param("id")
		id, err := strconv.Atoi(s)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		myPet, err := app.Pet.GetById(id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, struct {
			Data *entity.Pet `json:"data"`
		}{
			Data: myPet,
		})
	}
}

func DeletePet(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := ctx.Param("id")
		id, err := strconv.Atoi(s)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("ID: ", id)
		err = app.Pet.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok")
	}
}

func UpdatePet(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		s := ctx.Param("id")
		id, err := strconv.Atoi(s)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("ID: ", id)
		var pet entity.Pet
		pet.Id = id
		err = ctx.BindJSON(&pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		fmt.Println("UPDATE: ", pet)
		err = app.Pet.Update(pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok")
	}
}
