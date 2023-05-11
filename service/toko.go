package service

import (
	"shop/model"
	"shop/repo"
)

type Toko interface {
	GetByID(tokoid uint) (model.Toko, error)
	GetAll(limit uint, page uint) ([]model.Toko, error)

	Update(toko model.Toko) (model.Toko, error)
	Delete(tokoid uint) error
}

func NewServiceToko(repo repo.Toko) Toko {
	return &serviceToko{
		repo: repo,
	}
}

type serviceToko struct {
	repo repo.Toko
}

func (s *serviceToko) GetByID(tokoid uint) (model.Toko, error) {
	newToko, err := s.repo.Get(tokoid)
	if err != nil {
		return model.Toko{}, err
	}
	return newToko, nil
}

func (s *serviceToko) GetAll(limit uint, page uint) ([]model.Toko, error) {
	tokos, err := s.repo.GetAll(limit, page)
	if err != nil {
		return []model.Toko{}, err
	}
	return tokos, nil
}

func (s *serviceToko) Update(toko model.Toko) (model.Toko, error) {
	newToko, err := s.repo.Update(toko)
	if err != nil {
		return model.Toko{}, err
	}
	return newToko, nil
}

func (s *serviceToko) Delete(tokoid uint) error {
	err := s.repo.Delete(tokoid)
	if err != nil {
		return err
	}
	return nil
}
