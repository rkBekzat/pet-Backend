package server

import (
	"github.com/gorilla/websocket"
	"log"
	"pet/internal/entity"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	Message  chan *entity.Messages
	Id       int
	Username string
	RoomId   int
}

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func (c *Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}
		c.Conn.WriteJSON(message)
	}
}

func (c *Client) ReadMessage(hub *Hub) {
	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()
	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			return
		}

		msg := &entity.Messages{
			FromUsername: c.Username,
			FromUserId:   c.Id,
			CreatedAt:    time.Now().String(),
			Content:      string(m),
			RoomId:       c.RoomId,
		}
		hub.Broadcast <- msg
	}
}
