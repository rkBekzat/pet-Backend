package repository

import (
	"fmt"
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
	var id int
	query := "INSERT INTO pet (username, name) VALUES ($1, $2) RETURNING id"

	row := p.db.QueryRow(
		query,
		pet.Username,
		pet.Name,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (p *PetPostgres) Update(pet entity.Pet) error {
	query := "UPDATE pet SET name=$1 WHERE username=$2"
	exec, err := p.db.Exec(query, pet.Name, pet.Username)
	if err != nil {
		return err
	}
	fmt.Println(exec)
	return nil
}

func (p *PetPostgres) Delete(id int) error {
	return nil
}

func (p *PetPostgres) GetById(id int) (*entity.Pet, error) {
	var res *entity.Pet = new(entity.Pet)
	query := "SELECT * FROM pet WHERE id=$1"
	row := p.db.QueryRow(query, id)
	if err := row.Scan(
		&res.Id,
		&res.Username,
		&res.Species,
		&res.Breed,
		&res.Name,
		&res.DateOfBirth,
		&res.Color,
		&res.Sex,
		&res.Tattoo,
		&res.IssuedOrganization); err != nil {
		return nil, err
	}
	return res, nil
}

func (p *PetPostgres) GetAll() ([]*entity.Pet, error) {
	return nil, nil
}
