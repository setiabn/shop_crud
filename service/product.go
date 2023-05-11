package service

import (
	"shop/model"
	"shop/repo"
)

type Product interface {
	Create(product model.ProductComplete) (model.ProductComplete, error)
	Get(productcompleteid uint) (model.ProductComplete, error)
	GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.ProductComplete, error)
	Update(productcomplete model.ProductComplete) (model.ProductComplete, error)
	Delete(productcompleteid uint) error
}

func NewServiceProduct(repoProduct repo.Product, repoLogProduct repo.LogProduct, repoFotoProduct repo.FotoProduct) Product {
	return &serviceProduct{
		repoProduct:     repoProduct,
		repoLogProduct:  repoLogProduct,
		repoFotoProduct: repoFotoProduct,
	}
}

type serviceProduct struct {
	repoProduct     repo.Product
	repoLogProduct  repo.LogProduct
	repoFotoProduct repo.FotoProduct
}

func (s *serviceProduct) Create(product model.ProductComplete) (model.ProductComplete, error) {

	newProduct, err := s.repoProduct.Create(product.Product)
	if err != nil {
		return model.ProductComplete{}, err
	}

	newLogoProduct, err := s.repoLogProduct.Create(product.LogProduct)
	if err != nil {
		return model.ProductComplete{}, err
	}

	newFotoProduct, err := s.repoFotoProduct.Create(product.FotoProduct)
	if err != nil {
		return model.ProductComplete{}, err
	}

	return model.ProductComplete{Product: newProduct, LogProduct: newLogoProduct, FotoProduct: newFotoProduct}, nil
}

func (s *serviceProduct) Get(productid uint) (model.ProductComplete, error) {
	product, err := s.repoProduct.Get(productid)
	if err != nil {
		return model.ProductComplete{}, err
	}

	logProduct, fotoProduct, err := s.getProductDetail(productid)

	return model.ProductComplete{
		Product:     product,
		LogProduct:  logProduct,
		FotoProduct: fotoProduct,
	}, nil
}

func (s *serviceProduct) GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.ProductComplete, error) {
	var result []model.ProductComplete
	products, err := s.repoProduct.GetAll(limit, page, categoryid, tokoid)
	if err != nil {
		return []model.ProductComplete{}, err
	}

	for _, product := range products {
		logProduct, fotoProduct, err := s.getProductDetail(product.ID)
		if err != nil {
			return []model.ProductComplete{}, err
		}

		result = append(result, model.ProductComplete{
			Product:     product,
			LogProduct:  logProduct,
			FotoProduct: fotoProduct,
		})
	}

	return result, nil
}

func (s *serviceProduct) Update(product model.ProductComplete) (model.ProductComplete, error) {
	newProduct, err := s.repoProduct.Update(product.Product)
	if err != nil {
		return model.ProductComplete{}, err
	}

	return model.ProductComplete{
		Product: newProduct,
	}, nil
}

func (s *serviceProduct) Delete(productid uint) error {
	err := s.repoProduct.Delete(productid)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceProduct) getProductDetail(productid uint) (model.LogProduct, model.FotoProduct, error) {
	logProduct, err := s.repoLogProduct.GetByProductID(productid)
	if err != nil {
		return model.LogProduct{}, model.FotoProduct{}, err
	}

	fotoProduct, err := s.repoFotoProduct.GetByProductID(productid)
	if err != nil {
		return model.LogProduct{}, model.FotoProduct{}, err
	}

	return logProduct, fotoProduct, err

}
