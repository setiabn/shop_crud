package service

import (
	"shop/model"
	"shop/repo"
)

type Product interface {
	Create(product model.Product) (model.Product, error)
	Get(productcompleteid uint) (model.ProductComplete, error)
	GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.Product, error)
	Update(product model.Product) (model.Product, error)
	Delete(productcompleteid uint) error
}

func NewServiceProduct(repoProduct repo.Product, repoLogProduct repo.LogProduct, repoFotoProduct repo.FotoProduct, repoCategory repo.Category) Product {
	return &serviceProduct{
		repoProduct:     repoProduct,
		repoLogProduct:  repoLogProduct,
		repoFotoProduct: repoFotoProduct,
		repoCategory:    repoCategory,
	}
}

type serviceProduct struct {
	repoProduct     repo.Product
	repoLogProduct  repo.LogProduct
	repoFotoProduct repo.FotoProduct
	repoCategory    repo.Category
}

func (s *serviceProduct) Create(product model.Product) (model.Product, error) {

	category, err := s.repoCategory.Get(product.CategoryID)
	if err != nil {
		return model.Product{}, err
	}
	product.Category = category

	resultProduct, err := s.repoProduct.Create(product)
	if err != nil {
		return model.Product{}, err
	}

	return resultProduct, nil
}

func (s *serviceProduct) Get(productid uint) (model.ProductComplete, error) {
	product, err := s.repoProduct.Get(productid)
	if err != nil {
		return model.ProductComplete{}, err
	}

	logProduct, fotoProduct, err := s.getProductDetail(productid)
	if err != nil {
		return model.ProductComplete{}, err
	}

	return model.ProductComplete{
		Product:      product,
		LogProduct:   logProduct,
		FotoProducts: fotoProduct,
	}, nil
}

func (s *serviceProduct) GetAll(limit uint, page uint, categoryid uint, tokoid uint) ([]model.Product, error) {
	var result []model.Product
	products, err := s.repoProduct.GetAll(limit, page, categoryid, tokoid)
	if err != nil {
		return []model.Product{}, err
	}

	category, err := s.repoCategory.Get(categoryid)
	if err != nil {
		return []model.Product{}, err
	}

	for _, product := range products {
		logProduct, fotoProducts, err := s.getProductDetail(product.ID)
		if err != nil {
			return []model.Product{}, err
		}

		product.Category = category
		product.CategoryID = categoryid
		product.FotoProducts = fotoProducts
		product.TokoID = tokoid
		product.LogProduct = logProduct

		result = append(result, product)
	}

	return result, nil
}

func (s *serviceProduct) Update(product model.Product) (model.Product, error) {
	newProduct, err := s.repoProduct.Update(product)
	if err != nil {
		return model.Product{}, err
	}

	return newProduct, nil
}

func (s *serviceProduct) Delete(productid uint) error {
	err := s.repoProduct.Delete(productid)
	if err != nil {
		return err
	}
	return nil
}

func (s *serviceProduct) getProductDetail(productid uint) (model.LogProduct, []model.FotoProduct, error) {

	// logProduct, err := s.repoLogProduct.GetByProductID(productid)
	// if err != nil {
	// 	return model.LogProduct{}, model.FotoProduct{}, err
	// }

	fotoProducts, err := s.repoFotoProduct.GetByProductID(productid)
	if err != nil {
		return model.LogProduct{}, []model.FotoProduct{}, err
	}

	return model.LogProduct{}, fotoProducts, nil

}
