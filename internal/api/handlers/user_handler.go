package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/models"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/services"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "users": users})
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var payload models.CreateUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	id, err := h.service.CreateUser(&payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": false, "error": err})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": true, "user created with id": id})

}
