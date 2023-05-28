package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type Repository struct {
	Authorization
	Pets
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Pets:          NewPetPostgres(db),
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
