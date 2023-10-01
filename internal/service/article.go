package service

import (
	"pet/internal/entity"
	"pet/internal/repository"
)

type ArticleService struct {
	repo repository.Article
}

func NewArticleService(repo repository.Article) *ArticleService {
	return &ArticleService{repo: repo}
}

func (a *ArticleService) Create(post entity.Article, tag string) error {
	return nil
}

func (a *ArticleService) Update(post entity.Article) error {
	return nil
}

func (a *ArticleService) Delete(id int) error {
	return nil
}

func (a *ArticleService) GetById(id int) (*entity.Article, error) {
	return nil, nil
}

func (a *ArticleService) GetAll(username string) ([]*entity.Article, error) {
	return []*entity.Article{}, nil
}
