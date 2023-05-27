package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}

type Authorization interface {
	CreateUser(user entity.User) (int, error)
	GetUser(username, password string) (entity.User, error)
}

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}
func (auth *AuthPostgres) CreateUser(user entity.User) (int, error) {
	var id int
	query := "INSERT INTO users (name, username, password_hash) values ($1, $2, $3) RETURNING id"

	row := auth.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (auth *AuthPostgres) GetUser(username, password string) (entity.User, error) {
	var user entity.User
	query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"
	err := auth.db.Get(&user, query, username, password)

	return user, err
}
