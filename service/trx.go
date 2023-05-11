package service

import (
	"shop/model"
	"shop/repo"
)

type Trx interface {
	Create(transaction model.TrxComplete) (model.TrxComplete, error)
	GetByID(trxid uint) (model.TrxComplete, error)
	GetByUserID(userid uint) ([]model.TrxComplete, error)
}

func NewServiceTrx(repoTrx repo.Trx, repoDetaitTrx repo.DetailTrx) Trx {
	return &serviceTrx{
		rTrx:       repoTrx,
		rDetailTrx: repoDetaitTrx,
	}
}

type serviceTrx struct {
	rTrx       repo.Trx
	rDetailTrx repo.DetailTrx
}

func (s *serviceTrx) Create(transaction model.TrxComplete) (model.TrxComplete, error) {

	newTrx, err := s.rTrx.Create(transaction.Trx)
	if err != nil {
		return model.TrxComplete{}, err
	}
	newDetailTrx, err := s.rDetailTrx.Create(transaction.DetailTrx)
	if err != nil {
		return model.TrxComplete{}, err
	}
	return model.TrxComplete{Trx: newTrx, DetailTrx: newDetailTrx}, nil
}

func (s *serviceTrx) GetByID(trxid uint) (model.TrxComplete, error) {
	newTrx, err := s.rTrx.Get(trxid)
	if err != nil {
		return model.TrxComplete{}, err
	}

	newDetailTrx, err := s.rDetailTrx.GetByTrxID(newTrx.ID)
	if err != nil {
		return model.TrxComplete{}, err
	}

	return model.TrxComplete{Trx: newTrx, DetailTrx: newDetailTrx}, nil
}

func (s *serviceTrx) GetByUserID(userid uint) ([]model.TrxComplete, error) {

	var result []model.TrxComplete
	newTrxs, err := s.rTrx.GetByUserID(userid)
	if err != nil {
		return []model.TrxComplete{}, err
	}

	for _, trx := range newTrxs {
		newDetailTrx, err := s.rDetailTrx.GetByTrxID(trx.ID)
		if err != nil {
			return []model.TrxComplete{}, err
		}

		result = append(result, model.TrxComplete{Trx: trx, DetailTrx: newDetailTrx})
	}
	return result, nil
}
