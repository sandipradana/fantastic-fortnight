package repository

import (
	"fantastic-fortnight/backend/internal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(db *gorm.DB) ([]model.Product, error)
	GetByID(db *gorm.DB, id uint) (*model.Product, error)
	Create(db *gorm.DB, product *model.Product) error
	Update(db *gorm.DB, product *model.Product) error
	Delete(db *gorm.DB, id uint) error
}

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetAll(db *gorm.DB) ([]model.Product, error) {
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetByID(db *gorm.DB, id uint) (*model.Product, error) {
	var product model.Product
	if err := db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(db *gorm.DB, product *model.Product) error {
	if err := db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Update(db *gorm.DB, product *model.Product) error {
	if err := db.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func (r *productRepository) Delete(db *gorm.DB, id uint) error {
	if err := db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
