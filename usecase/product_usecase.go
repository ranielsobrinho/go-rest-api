package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (pu *ProductUsecase) GetProducts() ([]model.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUsecase) CreateProduct(product model.Product) (model.Product, error) {
	productId, err := pu.repository.CreateProduct(product)

	if err != nil {
		return model.Product{}, err
	}

	product.ID = productId

	return product, nil
}

func (pu *ProductUsecase) GetProductById(product_id int) (*model.Product, error) {
	product, err := pu.repository.GetProductById(product_id)

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (pu *ProductUsecase) DeleteProductById(product_id int) error {
	err := pu.repository.DeleteProductById(product_id)

	if err != nil {
		return err
	}

	return nil
}

func (pu *ProductUsecase) UpdateProduct(product_id int, update_product model.Product) (model.Product, error) {
	updatedProductId, err := pu.repository.UpdateProduct(product_id, update_product)
	if err != nil {
		return model.Product{}, err
	}

	update_product.ID = updatedProductId
	return update_product, nil
}
