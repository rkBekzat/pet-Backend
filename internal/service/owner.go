package service

import (
	"errors"
	"pet/internal/entity"
	"pet/internal/repository"
)

type Owner struct {
	repo repository.Owner
}

func NewOwnerService(repo repository.Owner) *Owner {
	return &Owner{repo: repo}
}

func (o *Owner) AddAddress(add entity.Address) error {
	if o.repo.HasAddress(add) {
		return errors.New("The address already added")
	}
	_, err := o.repo.AddAddress(add)
	if err != nil {
		return err
	}
	return nil
}

func (o *Owner) SetAddress(add entity.Address) error {
	if !o.repo.HasAddress(add) {
		return errors.New("The address not added")
	}
	_, err := o.repo.SetAddress(add)
	if err != nil {
		return err
	}
	return nil
}
