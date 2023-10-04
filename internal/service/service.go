package service

import (
	"pet/internal/entity"
	"pet/internal/repository"
)

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	GetName(id int) (string, error)
}

type PetAppInterface interface {
	Create(pet entity.Pet) error
	Update(pet entity.Pet) error
	Delete(id int) error
	GetById(id int) (*entity.Pet, error)
	GetAll(username string) ([]*entity.Pet, error)
}

type ArticleAppInterface interface {
	Create(post entity.Article, tag string) error
	Update(post entity.Article) error
	Delete(id int) error
	GetById(id int) (*entity.Article, error)
	GetAll(username string) ([]*entity.Article, error)
}

type OwnerAppInterface interface {
	AddAddress(add entity.Address) error
	SetAddress(add entity.Address) error
}

type MessageAppInterface interface {
	CreateRoom() int
	GetRooms(id int) []*entity.Room
	Save(message *entity.Messages)
	GetMessage(roomId int) []*entity.Messages
}

type Service struct {
	Auth    Authorization
	Pet     PetAppInterface
	Owner   OwnerAppInterface
	Message MessageAppInterface
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth:    NewAuthService(repo.Authorization),
		Pet:     NewPetService(repo.Pets),
		Owner:   NewOwnerService(repo.Owner),
		Message: NewMessages(repo.Message),
	}
}
