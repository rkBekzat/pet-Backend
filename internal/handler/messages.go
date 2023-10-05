package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"pet/internal/entity"
	"pet/internal/server"
	"pet/internal/service"
	"strconv"
)

func CreateRoom(app *service.Service, hub *server.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		otherId := ctx.Param("otherUser")
		id2, err := strconv.Atoi(otherId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		roomId, err := app.Message.CreateRoom(id, id2)
		if err != nil {
			ctx.JSON(http.StatusForbidden, err.Error())
			return
		}
		user1, _ := app.Auth.GetName(id)
		user2, _ := app.Auth.GetName(id2)
		room := server.NewRoom(user1 + user2)
		room.Id = roomId
		hub.Rooms[roomId] = room
		ctx.JSON(http.StatusOK, nil)
	}
}

func GetAllRooms(app *service.Service, hub *server.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		rooms, err := app.Message.GetRooms(id)
		if err != nil {
			return
		}
		ctx.JSON(http.StatusOK, rooms)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func JoinRoom(app *service.Service, hub *server.Hub) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := getUserId(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}

		pr := ctx.Param("room_id")
		roomId, err := strconv.Atoi(pr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		username, err := app.Auth.GetName(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		cl := &server.Client{
			Id:       id,
			Message:  make(chan *entity.Messages),
			Conn:     conn,
			RoomId:   roomId,
			Username: username,
		}

		hub.Register <- cl
		go cl.WriteMessage()
		cl.ReadMessage(hub)
	}
}
