package model

type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type City struct {
	ID         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}
