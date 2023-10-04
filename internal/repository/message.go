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

func (m *MessagePostgres) Save(message *entity.Messages) {

}

func (m *MessagePostgres) GetMessages(roomId int) []*entity.Messages {

	return []*entity.Messages{}
}
