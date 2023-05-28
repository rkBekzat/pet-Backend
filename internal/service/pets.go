package service

import (
	"pet/internal/entity"
	"pet/internal/repository"
)

type PetService struct {
	repo repository.Pets
}

func NewPetService(repo repository.Pets) *PetService {
	return &PetService{
		repo: repo,
	}
}

func (p *PetService) Create(pet entity.Pet) error {
	return nil
}

func (p *PetService) Update(pet entity.Pet) error {
	return nil
}

func (p *PetService) Delete(id int) error {
	return nil
}

func (p *PetService) GetById(id int) (*entity.Pet, error) {
	return nil, nil
}

func (p *PetService) GetAll() ([]*entity.Pet, error) {
	return nil, nil
}
