package repo

import (
	"errors"
	"shop/model"

	"gorm.io/gorm"
)

type DetailTrx interface {
	Create(detailtrx model.DetailTrx) (model.DetailTrx, error)
	Get(detailtrxid uint) (model.DetailTrx, error)
	GetByTrxID(id uint) (model.DetailTrx, error)
	GetByLogProductID(id uint) ([]model.DetailTrx, error)
	GetByTokoID(id uint) ([]model.DetailTrx, error)

	Update(detailtrx model.DetailTrx) (model.DetailTrx, error)
	Delete(detailtrxid uint) error
}

func NewDetailTrxRepo(db *gorm.DB) DetailTrx {
	return &repodetailtrx{
		DB: db,
	}
}

type repodetailtrx struct {
	DB *gorm.DB
}

func (r *repodetailtrx) Create(detailtrx model.DetailTrx) (model.DetailTrx, error) {

	result := r.DB.Create(&detailtrx)
	if result.Error != nil {
		return model.DetailTrx{}, result.Error
	}
	if result.RowsAffected == 0 {
		return model.DetailTrx{}, errors.New("not created")
	}

	return detailtrx, nil
}

func (r *repodetailtrx) Get(id uint) (model.DetailTrx, error) {
	var detailtrx model.DetailTrx
	result := r.DB.First(&detailtrx, id)
	if result.Error != nil {
		return model.DetailTrx{}, result.Error
	}

	return detailtrx, nil
}

func (r *repodetailtrx) GetByTrxID(id uint) (model.DetailTrx, error) {
	var detailtrx model.DetailTrx

	result := r.DB.Where(&model.DetailTrx{TrxID: id}).First(&detailtrx)
	if result.Error != nil {
		return model.DetailTrx{}, result.Error
	}
	return detailtrx, nil
}

func (r *repodetailtrx) GetByLogProductID(id uint) ([]model.DetailTrx, error) {
	var detailtrxs []model.DetailTrx
	result := r.DB.Where(&model.DetailTrx{}).Find(&detailtrxs)
	if result.Error != nil {
		return []model.DetailTrx{}, result.Error
	}
	return detailtrxs, nil
}

func (r *repodetailtrx) GetByTokoID(id uint) ([]model.DetailTrx, error) {
	var detailtrxs []model.DetailTrx
	result := r.DB.Where(&model.DetailTrx{TokoID: id}).Find(&detailtrxs)
	if result.Error != nil {
		return []model.DetailTrx{}, result.Error
	}
	return detailtrxs, nil
}

func (r *repodetailtrx) Update(detailtrx model.DetailTrx) (model.DetailTrx, error) {

	var olddata model.DetailTrx

	result := r.DB.First(&olddata, detailtrx.ID)
	if result.Error != nil {
		return model.DetailTrx{}, result.Error
	}

	detailtrx.CreatedAt = olddata.CreatedAt

	result = r.DB.Create(&detailtrx)
	if result.Error != nil {
		return model.DetailTrx{}, result.Error
	}

	return detailtrx, nil
}

func (r *repodetailtrx) Delete(id uint) error {

	result := r.DB.Delete(&model.DetailTrx{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
