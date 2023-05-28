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
	var id int
	query := "INSERT INTO pet (username, species, breed,  name, date_of_birth, color, sex, tattoo, issued_organization) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"

	row := p.db.QueryRow(
		query,
		pet.Username,
		pet.Species,
		pet.Breed,
		pet.Name,
		pet.DateOfBirth,
		pet.Color,
		pet.Sex,
		pet.Tattoo,
		pet.IssuedOrganization,
	)
	if err := row.Scan(&id); err != nil {
		return err
	}
	return nil
}

func (p *PetPostgres) Update(pet entity.Pet) error {
	query := "UPDATE pet SET name=$1, species=$3, breed=$4, date_of_birth=$5, color=$6, sex=$7, tattoo=$8, issued_organization=$9  WHERE id=$2"
	_, err := p.db.Exec(query, pet.Name, pet.Id, pet.Species, pet.Breed, pet.DateOfBirth, pet.Color, pet.Sex, pet.Tattoo, pet.IssuedOrganization)
	if err != nil {
		return err
	}
	return nil
}

func (p *PetPostgres) Delete(id int) error {
	query := "DELETE FROM pet WHERE id=$1"
	_, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
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
	query := "SELECT * FROM pet"
	row, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	var result []*entity.Pet
	for row.Next() {
		var res *entity.Pet = new(entity.Pet)
		if err = row.Scan(
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
		result = append(result, res)
	}
	return result, nil
}
