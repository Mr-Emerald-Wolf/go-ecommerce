package interfaces

import "github.com/mr-emerald-wolf/go-ecommerce/internal/models"

type UserRepository interface {
	GetUserByID(id int) (*models.User, error)
	CreateUser(user *models.User) (int, error)
	UpdateUser(user *models.User) error
	DeleteUser(id int) error
	GetAllUsers() ([]models.User, error)
}
