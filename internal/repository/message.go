package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type MessagePostgres struct {
	db *sqlx.DB
}

func NewMessage(db *sqlx.DB) *MessagePostgres {
	return &MessagePostgres{db: db}
}

// create room and return id of room
func (m *MessagePostgres) CreateRoom() int {
	return 0
}

func (m *MessagePostgres) GetRooms(id int) []*entity.Room {
	return []*entity.Room{}
}

func (m *MessagePostgres) Save(message *entity.Messages) {

}

func (m *MessagePostgres) GetMessages(roomId int) []*entity.Messages {

	return []*entity.Messages{}
}
