package server

import (
	"github.com/gorilla/websocket"
	"pet/internal/entity"
	"time"
)

type Client struct {
	Conn     *websocket.Conn
	Message  entity.Messages
	Id       int
	Username string
}

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

func (c *Client) WriteMessage() {
	ticker := time.NewTimer(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {}
		//_, payload, err := c.Conn.ReadMessage()
		//if err != nil {
		//	return
		//}
	}
}
