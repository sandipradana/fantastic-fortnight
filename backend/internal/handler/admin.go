package handler

import (
	"fantastic-fortnight/backend/internal/model"
	"fantastic-fortnight/backend/internal/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AdminHandler struct {
	adminService service.AdminService
}

func NewAdminHandler(service service.AdminService) *AdminHandler {
	return &AdminHandler{adminService: service}
}

func (h *AdminHandler) GetAll(c *fiber.Ctx) error {
	admins, err := h.adminService.GetAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admins)
}

func (h *AdminHandler) Get(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	admin, err := h.adminService.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) Create(c *fiber.Ctx) error {
	admin := new(model.Admin)
	if err := c.BodyParser(admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.adminService.Create(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	admin, err := h.adminService.GetByID(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Admin not found"})
	}

	if err := c.BodyParser(admin); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.adminService.Update(admin); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(admin)
}

func (h *AdminHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.adminService.Delete(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *AdminHandler) Login(c *fiber.Ctx) error {
	loginData := model.AdminLogin{}
	if err := c.BodyParser(&loginData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	token, err := h.adminService.Login(loginData)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"token": token})
}
