package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type OwnerPostgres struct {
	db *sqlx.DB
}

func NewOwnerPostgres(db *sqlx.DB) *OwnerPostgres {
	return &OwnerPostgres{
		db: db,
	}
}

func (o *OwnerPostgres) AddAddress(add entity.Address) (*entity.Address, error) {
	query := "INSERT INTO Address (country, city, state_province, zip_postal_code, address_line) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	var id int
	row := o.db.QueryRow(query, add.Country, add.City, add.StateProvince, add.ZipPostalCode, add.AddressLine)
	if err := row.Scan(&id); err != nil {
		return nil, err
	}
	add.AddressId = id
	return &add, nil
}

func (o *OwnerPostgres) HasAddress(add entity.Address) bool {
	query := "SELECT count(*) FROM Address WHERE address_id=$1"
	var cnt int
	row := o.db.QueryRow(query, add.AddressId)
	if err := row.Scan(&cnt); err != nil {
		return false
	}
	return cnt > 0
}

func (o *OwnerPostgres) SetAddress(add entity.Address) (*entity.Address, error) {
	query := "UPDATE Address SET country=$2, city=$3, state_province=$4, zip_postal_code=$5, address_line=$6  HERE address_id=$1"
	_, err := o.db.Exec(query, add.AddressId, add.Country, add.City, add.StateProvince, add.ZipPostalCode, add.ZipPostalCode)
	if err != nil {
		return nil, err
	}
	return &add, nil
}

func (o *OwnerPostgres) Found() {

}

func (o *OwnerPostgres) Lost() {

}
