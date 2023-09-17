package services

import (
	"github.com/mr-emerald-wolf/go-ecommerce/internal/interfaces"
	"github.com/mr-emerald-wolf/go-ecommerce/internal/models"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(repo interfaces.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.userRepository.GetUserByID(id)
}

func (s *UserService) CreateUser(user *models.User) (int, error) {
	return s.userRepository.CreateUser(user)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepository.UpdateUser(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.userRepository.DeleteUser(id)
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepository.GetAllUsers()
}
