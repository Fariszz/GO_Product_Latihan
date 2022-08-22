package product

import (
	"GoProductLatihan/database"
	"GoProductLatihan/models"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]models.Product, error)
	FindByID(id int) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(product models.Product) (models.Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{
		db: database.DbManager(),
	}
}

func ( r *repository) FindAll() ([]models.Product, error) {
	products := []models.Product{}

	err := r.db.Find(&products).Error

	if err != nil {
		return products, err
	}

	return products, nil
}

func (r *repository) FindByID(id int) (models.Product, error) {
	var product models.Product

	err := r.db.First(&product, id).Error

	if err != nil {
		return product, err
	}

	return product, nil

}

func (r *repository) Create(product models.Product) (models.Product, error) {
	err := r.db.Create(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
	
}

func (r *repository) Update(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func(r *repository) Delete(product models.Product)(models.Product, error){
	err := r.db.Delete(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}