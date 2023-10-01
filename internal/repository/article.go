package repository

import (
	"github.com/jmoiron/sqlx"
	"pet/internal/entity"
)

type ArticlePostgres struct {
	db *sqlx.DB
}

func NewArticlePostgres(db *sqlx.DB) *ArticlePostgres {
	return &ArticlePostgres{db: db}
}

func (a *ArticlePostgres) Create(post entity.Article) error {

	return nil
}
func (a *ArticlePostgres) Update(post entity.Article) error {
	return nil
}

func (a *ArticlePostgres) Delete(id int) error {
	return nil
}

func (a *ArticlePostgres) GetById(id int) (*entity.Article, error) {
	return nil, nil
}

func (a *ArticlePostgres) GetAll(username string) ([]*entity.Article, error) {
	return []*entity.Article{}, nil
}
