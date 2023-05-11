package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"shop/model"
)

type ProvCity interface {
	GetAllProvincies() ([]model.Province, error)
	GetAllCities(provId string) ([]model.City, error)
	GetDetaiProvince(provId string) (model.Province, error)
	GetDetailCity(cityId string) (model.City, error)
}

func NewProvCityRepo() ProvCity {
	return &repoprovcity{}

}

type repoprovcity struct{}

func (r *repoprovcity) GetAllProvincies() ([]model.Province, error) {

	url := "https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json"
	body, err := fetchBody(url)
	if err != nil {
		return []model.Province{}, err
	}

	var result []model.Province
	if err := json.Unmarshal(body, &result); err != nil {
		return []model.Province{}, err
	}

	return result, nil
}

func (r *repoprovcity) GetAllCities(provId string) ([]model.City, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%v.json", provId)
	body, err := fetchBody(url)
	if err != nil {
		return []model.City{}, err
	}

	var result []model.City
	if err := json.Unmarshal(body, &result); err != nil {
		return []model.City{}, err

	}
	return result, nil
}

func (r *repoprovcity) GetDetaiProvince(provId string) (model.Province, error) {
	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/province/%v.json", provId)
	body, err := fetchBody(url)
	if err != nil {
		return model.Province{}, err
	}

	var result model.Province
	if err := json.Unmarshal(body, &result); err != nil {
		return model.Province{}, err
	}
	return result, nil
}

func (r *repoprovcity) GetDetailCity(cityId string) (model.City, error) {

	url := fmt.Sprintf("https://emsifa.github.io/api-wilayah-indonesia/api/regency/%v.json", cityId)
	body, err := fetchBody(url)
	if err != nil {
		return model.City{}, err
	}

	var result model.City
	if err := json.Unmarshal(body, &result); err != nil {
		return model.City{}, err
	}

	return result, nil
}

func fetchBody(url string) ([]byte, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
