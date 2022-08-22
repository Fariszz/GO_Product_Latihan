package product

import "GoProductLatihan/models"

type Service interface {
	GetProducts() ([]models.Product, error)
	GetProductById(input GetProductDetailInput) (models.Product, error)
	CreateProduct(input CreateProductInput) (models.Product, error)
	UpdateProduct(input GetProductDetailInput, inputData CreateProductInput) (models.Product, error)
	DeleteProduct(input GetProductDetailInput) (models.Product, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{
		repository,
	}
}

func (s *service) GetProducts() ([]models.Product, error){
	products, err := s.repository.FindAll()
	
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetProductById(input GetProductDetailInput) (models.Product, error){
	product, err := s.repository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (s *service) CreateProduct(input CreateProductInput) (models.Product, error) {
	product := models.Product{}
	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.Description = input.Description	

	newProduct, err := s.repository.Create(product)

	if err != nil {
		return newProduct, err
	}

	return newProduct, nil
}

func(s *service) UpdateProduct(input GetProductDetailInput, inputData CreateProductInput) (models.Product, error){
	product, err := s.repository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	product.Name = inputData.Name
	product.Price = inputData.Price
	product.Stock = inputData.Stock
	product.Description = inputData.Description

	updatedProduct, err := s.repository.Update(product)

	if err != nil {
		return product, err
	}

	return updatedProduct, nil
}

func (s *service) DeleteProduct(input GetProductDetailInput) (models.Product, error){
	product, err := s.repository.FindByID(input.ID)

	if err != nil {
		return product, err
	}

	deletedProduct, err := s.repository.Delete(product)

	if err != nil {
		return product, err
	}

	return deletedProduct, nil
	
}