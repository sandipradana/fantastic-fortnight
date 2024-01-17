package service

import (
	"fantastic-fortnight/backend/internal/model"
	"fantastic-fortnight/backend/internal/repository"

	"gorm.io/gorm"
)

type ProductService interface {
	GetAll() ([]model.Product, error)
	GetByID(id uint) (*model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
}

type productService struct {
	productRepo repository.ProductRepository
	db          *gorm.DB
}

func NewProductService(db *gorm.DB, repo repository.ProductRepository) ProductService {
	return &productService{productRepo: repo, db: db}
}

func (s *productService) GetAll() ([]model.Product, error) {
	return s.productRepo.GetAll(s.db)
}

func (s *productService) GetByID(id uint) (*model.Product, error) {
	return s.productRepo.GetByID(s.db, id)
}

func (s *productService) Create(product *model.Product) error {
	return s.productRepo.Create(s.db, product)
}

func (s *productService) Update(product *model.Product) error {
	return s.productRepo.Update(s.db, product)
}

func (s *productService) Delete(id uint) error {
	return s.productRepo.Delete(s.db, id)
}
