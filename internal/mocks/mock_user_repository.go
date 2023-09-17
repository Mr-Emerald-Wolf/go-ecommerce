package mocks

import (
	"fmt"

	"github.com/mr-emerald-wolf/go-ecommerce/internal/models"
)

// MockUserRepository is a mock implementation of the UserRepository interface.
type MockUserRepository struct {
	users  []models.User
	nextID int
}

// NewMockUserRepository creates a new instance of the MockUserRepository.
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users:  []models.User{},
		nextID: 1,
	}
}

// GetUserByID retrieves a user by ID from the mock repository.
func (r *MockUserRepository) GetUserByID(id int) (*models.User, error) {
	for _, user := range r.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// CreateUser adds a new user to the mock repository.
func (r *MockUserRepository) CreateUser(user *models.User) (int, error) {
	user.ID = r.nextID
	r.nextID++
	r.users = append(r.users, *user)
	return user.ID, nil
}

// UpdateUser updates a user in the mock repository.
func (r *MockUserRepository) UpdateUser(user *models.User) error {
	for i, existingUser := range r.users {
		if existingUser.ID == user.ID {
			r.users[i] = *user
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", user.ID)
}

// DeleteUser deletes a user from the mock repository by ID.
func (r *MockUserRepository) DeleteUser(id int) error {
	for i, user := range r.users {
		if user.ID == id {
			// Remove the user from the slice
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with ID %d not found", id)
}

// GetAllUsers retrieves all users from the mock repository.
func (r *MockUserRepository) GetAllUsers() ([]models.User, error) {
	return r.users, nil
}
