package vercel

import (
	"fantastic-fortnight/backend/internal/handler"
	"fantastic-fortnight/backend/internal/repository"
	"fantastic-fortnight/backend/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var router *fiber.App
var routerAdaptor http.HandlerFunc

func init() {

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

	routerGroup.Get("/admins", adminHandler.GetAll)
	routerGroup.Post("/admins/login", adminHandler.Login)
	routerGroup.Get("/admins/:id", adminHandler.Get)
	routerGroup.Post("/admins", adminHandler.Create)
	routerGroup.Put("/admins/:id", adminHandler.Update)
	routerGroup.Delete("/admins/:id", adminHandler.Delete)

	routerGroup.Get("/products", productHandler.GetAll)
	routerGroup.Get("/products/:id", productHandler.Get)
	routerGroup.Post("/products", productHandler.Create)
	routerGroup.Put("/products/:id", productHandler.Update)
	routerGroup.Delete("/products/:id", productHandler.Delete)

	routerAdaptor = adaptor.FiberApp(router)
}

func Handler(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	routerAdaptor.ServeHTTP(httpResponse, httpRequest)
}
