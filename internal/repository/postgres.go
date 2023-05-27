package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

//ALTER TABLE Articles
//ADD CONSTRAINT fk_articles_owner
//FOREIGN KEY (username) REFERENCES Owner(username)
//ON DELETE CASCADE;
//
//ALTER TABLE Articles
//ADD CONSTRAINT fk_articles_pet
//FOREIGN KEY (id) REFERENCES Pet(id)
//ON DELETE CASCADE;
//
//ALTER TABLE Owner
//ADD CONSTRAINT fk_owner_address
//FOREIGN KEY (address_id) REFERENCES Address(address_id)
//ON DELETE CASCADE;
//
//ALTER TABLE Pet
//ADD CONSTRAINT fk_pet_owner
//FOREIGN KEY (username) REFERENCES Owner(username)
//ON DELETE CASCADE;
//
//ALTER TABLE Found
//ADD CONSTRAINT fk_found_owner
//FOREIGN KEY (username) REFERENCES Owner(username)
//ON DELETE CASCADE;
//
//ALTER TABLE Found
//ADD CONSTRAINT fk_found_address
//FOREIGN KEY (address_id) REFERENCES Address(address_id)
//ON DELETE CASCADE;
//
//ALTER TABLE Found
//ADD CONSTRAINT fk_found_pet
//FOREIGN KEY (id) REFERENCES Pet(id)
//ON DELETE CASCADE;
//
//ALTER TABLE Lost
//ADD CONSTRAINT fk_lost_owner
//FOREIGN KEY (username) REFERENCES Owner(username)
//ON DELETE CASCADE;
//
//ALTER TABLE Lost
//ADD CONSTRAINT fk_lost_address
//FOREIGN KEY (address_id) REFERENCES Address(address_id)
//ON DELETE CASCADE;
//
//ALTER TABLE Lost
//ADD CONSTRAINT fk_lost_pet
//FOREIGN KEY (id) REFERENCES Pet(id)
//ON DELETE CASCADE;
