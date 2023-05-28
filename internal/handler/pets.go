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
		_, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		var pet entity.Pet
		err = ctx.BindJSON(&pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
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
		all, err := app.Pet.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, all)
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
		fmt.Println("ID: ", id)
		myPet, err := app.Pet.GetById(id)
		fmt.Println("PET:? ", myPet)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		ctx.JSON(http.StatusOK, myPet)
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
		err = ctx.BindJSON(&pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err = app.Pet.Update(pet)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, "ok")
	}
}
