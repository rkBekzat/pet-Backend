package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"pet/internal/service"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func UserIdentity(app *service.Service) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(authorizationHeader)
		if header == "" {
			ctx.JSON(http.StatusUnauthorized, "empty auth head")
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			ctx.JSON(http.StatusUnauthorized, "invalid auth")
			return
		}

		userId, err := app.Auth.ParseToken(headerParts[1])
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Set(userCtx, userId)
	}
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
