package service

import (
	"fantastic-fortnight/backend/internal/model"
	"fantastic-fortnight/backend/internal/repository"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AdminService interface {
	GetAll() ([]model.Admin, error)
	GetByID(id uint) (*model.Admin, error)
	Create(admin *model.Admin) error
	Update(admin *model.Admin) error
	Delete(id uint) error
	Login(adminLogin model.AdminLogin) (*string, error)
}

type adminService struct {
	adminRepo repository.AdminRepository
	db        *gorm.DB
}

func NewAdminService(db *gorm.DB, repo repository.AdminRepository) AdminService {
	return &adminService{db: db, adminRepo: repo}
}

func (s *adminService) GetAll() ([]model.Admin, error) {
	return s.adminRepo.GetAll(s.db)
}

func (s *adminService) GetByID(id uint) (*model.Admin, error) {
	return s.adminRepo.GetByID(s.db, id)
}

func (s *adminService) Create(admin *model.Admin) error {
	passwordEncoded, _ := bcrypt.GenerateFromPassword([]byte(admin.Password), 14)
	admin.Password = string(passwordEncoded)
	return s.adminRepo.Create(s.db, admin)
}

func (s *adminService) Update(admin *model.Admin) error {
	return s.adminRepo.Update(s.db, admin)
}

func (s *adminService) Delete(id uint) error {
	return s.adminRepo.Delete(s.db, id)
}

func (s *adminService) Login(adminLogin model.AdminLogin) (*string, error) {

	admin, err := s.adminRepo.GetAdminByEmail(s.db, adminLogin.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminLogin.Password)); err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid email or password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["admin_id"] = admin.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("APP_SECRET")))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Failed to generate JWT token")
	}

	return &tokenString, nil
}
