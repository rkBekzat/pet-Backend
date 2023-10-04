package server

import (
	"pet/internal/entity"
)

type Room struct {
	Id      int             `json:"room_id"`
	Name    string          `json:"room_name"`
	Clients map[int]*Client `json:"clients"`
}

type Hub struct {
	Rooms      map[int]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *entity.Messages
}

func NewRoom(name string) *Room {
	return &Room{
		Name:    name,
		Clients: make(map[int]*Client),
	}
}

func NewHub() *Hub {
	return &Hub{
		Rooms:      make(map[int]*Room),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *entity.Messages),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				r := h.Rooms[cl.RoomId]
				r.Clients[cl.Id] = cl
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomId]; ok {
				delete(h.Rooms[cl.RoomId].Clients, cl.Id)
				close(cl.Message)
			}
		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomId]; ok {
				for _, cl := range h.Rooms[m.RoomId].Clients {
					cl.Message <- m
				}
			}
		}
	}
}
