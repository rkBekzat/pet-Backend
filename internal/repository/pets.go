package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type PetPostgres struct {
	db *sqlx.DB
}

func NewPetPostgres(db *sqlx.DB) *PetPostgres {
	return &PetPostgres{
		db: db,
	}
}

func (p *PetPostgres) Create(pet entity.Pet) error {
	return nil
}

func (p *PetPostgres) Update(pet entity.Pet) error {
	return nil
}

func (p *PetPostgres) Delete(id int) error {
	return nil
}

func (p *PetPostgres) GetById(id int) (*entity.Pet, error) {
	return nil, nil
}

func (p *PetPostgres) GetAll() ([]*entity.Pet, error) {
	return nil, nil
}
