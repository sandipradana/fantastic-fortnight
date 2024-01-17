package main

import (
	"fantastic-fortnight/backend/internal/handler"
	"fantastic-fortnight/backend/internal/middleware"
	"fantastic-fortnight/backend/internal/repository"
	"fantastic-fortnight/backend/internal/service"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var router *fiber.App

func init() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_TIMEZONE"))
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	adminRepo := repository.NewAdminRepository()
	adminService := service.NewAdminService(db, adminRepo)
	adminHandler := handler.NewAdminHandler(adminService)

	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(db, productRepo)
	productHandler := handler.NewProductHandler(productService)

	router = fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			errFiber, ok := err.(*fiber.Error)
			if ok {
				return c.Status(errFiber.Code).JSON(fiber.Map{"error": errFiber.Message})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		},
	})

	routerGroup := router.Group("/api")

	routerGroup.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World !")
	})

	routerGroup.Get("/admins", middleware.Auth(adminService), adminHandler.GetAll)
	routerGroup.Post("/admins/login", adminHandler.Login)
	routerGroup.Get("/admins/:id", middleware.Auth(adminService), adminHandler.Get)
	routerGroup.Post("/admins", middleware.Auth(adminService), adminHandler.Create)
	routerGroup.Put("/admins/:id", middleware.Auth(adminService), adminHandler.Update)
	routerGroup.Delete("/admins/:id", middleware.Auth(adminService), adminHandler.Delete)

	routerGroup.Get("/products", middleware.Auth(adminService), productHandler.GetAll)
	routerGroup.Get("/products/:id", middleware.Auth(adminService), productHandler.Get)
	routerGroup.Post("/products", middleware.Auth(adminService), productHandler.Create)
	routerGroup.Put("/products/:id", middleware.Auth(adminService), productHandler.Update)
	routerGroup.Delete("/products/:id", middleware.Auth(adminService), productHandler.Delete)
}

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8000"
	}
	err := router.Listen(":" + port)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
