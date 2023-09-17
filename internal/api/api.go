package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/api/handlers"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/mocks"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/services"
)

// Create user routes
func CreateUserRoutes(incomingRoutes *fiber.App) {
	MockUserRepository := mocks.NewMockUserRepository()
	UserService := services.NewUserService(MockUserRepository)
	UserHandler := handlers.NewUserHandler(*UserService)
	userGroup := incomingRoutes.Group("/users")
	userGroup.Get("/all", UserHandler.GetUsers)
	userGroup.Post("/create", UserHandler.CreateUser)
}
