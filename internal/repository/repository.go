package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"pet/internal/entity"
)

type Repository struct {
	Authorization
	Pets
	Owner
	Message
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Pets:          NewPetPostgres(db),
		Message:       NewMessage(db),
	}
}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
	GetNameById(id int) (string, error)
}

type Pets interface {
	Create(pet entity.Pet) error
	Update(pet entity.Pet) error
	Delete(id int) error
	GetById(id int) (*entity.Pet, error)
	GetAll(username string) ([]*entity.Pet, error)
}

type Article interface {
	Create(post entity.Article) error
	Update(post entity.Article) error
	Delete(id int) error
	GetById(id int) (*entity.Article, error)
	GetAll(username string) ([]*entity.Article, error)
}

type Owner interface {
	AddAddress(address entity.Address) (*entity.Address, error)
	SetAddress(address entity.Address) (*entity.Address, error)
	HasAddress(add entity.Address) bool
	Found()
	Lost()
}

type Message interface {
	CreateRoom() int
	GetRooms(id int) []*entity.Room
	Save(message *entity.Messages)
	GetMessages(roomId int) []*entity.Messages
}
