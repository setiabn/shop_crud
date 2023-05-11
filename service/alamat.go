package service

import (
	"shop/model"
	"shop/repo"
)

type Alamat interface {
	Create(alamat model.Alamat) (model.Alamat, error)
	GetByID(alamatid uint) (model.Alamat, error)
	GetByUserID(userid uint) ([]model.Alamat, error)
	Update(alamat model.Alamat) (model.Alamat, error)
	Delete(alamat model.Alamat) error
}

func NewServiceAlamat(repo repo.Alamat) Alamat {
	return &serviceAlamat{
		repo: repo,
	}
}

type serviceAlamat struct {
	repo repo.Alamat
}

func (s *serviceAlamat) Create(alamat model.Alamat) (model.Alamat, error) {

	newAlamat, err := s.repo.Create(alamat)
	if err != nil {
		return model.Alamat{}, err
	}
	return newAlamat, nil
}

func (s *serviceAlamat) GetByID(alamatid uint) (model.Alamat, error) {
	newAlamat, err := s.repo.Get(alamatid)
	if err != nil {
		return model.Alamat{}, err
	}
	return newAlamat, nil
}

func (s *serviceAlamat) GetByUserID(userid uint) ([]model.Alamat, error) {
	allAlamat, err := s.repo.GetByUserID(userid)
	if err != nil {
		return []model.Alamat{}, err
	}
	return allAlamat, nil
}

func (s *serviceAlamat) Update(alamat model.Alamat) (model.Alamat, error) {
	newAlamat, err := s.repo.Update(alamat)
	if err != nil {
		return model.Alamat{}, err
	}
	return newAlamat, nil
}

func (s *serviceAlamat) Delete(alamat model.Alamat) error {
	err := s.repo.Delete(alamat)
	if err != nil {
		return err
	}
	return nil
}
