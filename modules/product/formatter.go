package product

import "GoProductLatihan/models"

type ProductFormatter struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Stock       int     `json:"stock"`
}

func FormatProduct(product models.Product) ProductFormatter {
	productFormatter := ProductFormatter{}
	productFormatter.ID = product.ID
	productFormatter.Name = product.Name
	productFormatter.Price = product.Price
	productFormatter.Description = product.Description
	productFormatter.Stock = product.Stock

	return productFormatter
}

func FormatProducts(products []models.Product) []ProductFormatter  {
	productsFormatter := []ProductFormatter{}

	for _, product := range products {
		ProductFormatter := FormatProduct(product)
		productsFormatter = append(productsFormatter, ProductFormatter)
	}

	return productsFormatter
}