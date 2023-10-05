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

func (m *MessageService) CreateRoom(user1, user2 int) (int, error) {
	return m.repo.CreateRoom(user1, user2)
}

func (m *MessageService) GetRooms(id int) ([]*entity.Room, error) {
	return m.repo.GetRooms(id)
}

func (m *MessageService) Save(message *entity.Messages) error {
	return m.repo.Save(message)
}

func (m *MessageService) GetMessage(roomID int) ([]*entity.Messages, error) {
	return m.repo.GetMessages(roomID)
}
