package service

import (
	"pet/internal/entity"
	"pet/internal/repository"
)

type MessageService struct {
	repo repository.Message
}

func NewMessages(repo repository.Message) *MessageService {
	return &MessageService{
		repo: repo,
	}
}

func (m *MessageService) CreateRoom() int {
	return m.repo.CreateRoom()
}

func (m *MessageService) GetRooms(id int) []*entity.Room {
	return m.repo.GetRooms(id)
}

func (m *MessageService) Save(message *entity.Messages) {
	m.repo.Save(message)
}

func (m *MessageService) GetMessage(roomID int) []*entity.Messages {
	return m.repo.GetMessages(roomID)
}
