package repository

import (
	"fantastic-fortnight/backend/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAll(db *gorm.DB) ([]model.Admin, error)
	GetByID(db *gorm.DB, id uint) (*model.Admin, error)
	Create(db *gorm.DB, admin *model.Admin) error
	Update(db *gorm.DB, admin *model.Admin) error
	Delete(db *gorm.DB, id uint) error
	GetAdminByEmail(db *gorm.DB, email string) (*model.Admin, error)
}

type adminRepository struct {
}

func NewAdminRepository() AdminRepository {
	return &adminRepository{}
}

func (r *adminRepository) GetAll(db *gorm.DB) ([]model.Admin, error) {
	var items []model.Admin
	if err := db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *adminRepository) GetByID(db *gorm.DB, id uint) (*model.Admin, error) {
	var item model.Admin
	if err := db.First(&item, id).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *adminRepository) Create(db *gorm.DB, admin *model.Admin) error {
	if err := db.Create(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) Update(db *gorm.DB, admin *model.Admin) error {
	if err := db.Save(admin).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) Delete(db *gorm.DB, id uint) error {
	if err := db.Delete(&model.Admin{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *adminRepository) GetAdminByEmail(db *gorm.DB, email string) (*model.Admin, error) {
	var admin model.Admin
	if err := db.Where("email = ?", email).First(&admin).Error; err != nil {
		return nil, err
	}
	return &admin, nil
}
