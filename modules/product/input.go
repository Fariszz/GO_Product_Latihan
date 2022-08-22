package product

type GetProductDetailInput struct{
	ID int  `uri:"id" json:"id" binding:"required"`
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Price       float64    `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
}
