package service

import (
	"shop/model"
	"shop/repo"
)

type ProvCity interface {
	GetAllProvincies() ([]model.Province, error)
	GetAllCities(provId string) ([]model.City, error)
	GetDetaiProvince(provId string) (model.Province, error)
	GetDetailCity(cityId string) (model.City, error)
}

func NewServiceProvCity(repo repo.ProvCity) ProvCity {
	return &serviceProvCity{
		repo: repo,
	}
}

type serviceProvCity struct {
	repo repo.ProvCity
}

func (s *serviceProvCity) GetAllProvincies() ([]model.Province, error) {
	return s.repo.GetAllProvincies()
}

func (s *serviceProvCity) GetAllCities(provId string) ([]model.City, error) {
	return s.repo.GetAllCities(provId)
}

func (s *serviceProvCity) GetDetaiProvince(provId string) (model.Province, error) {
	return s.repo.GetDetaiProvince(provId)
}

func (s *serviceProvCity) GetDetailCity(cityId string) (model.City, error) {
	return s.repo.GetDetailCity(cityId)
}
