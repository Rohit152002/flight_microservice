package services

import (
	"errors"

	"user-services/models"
	"user-services/repositories"

	"go.uber.org/zap"
)

type UserServices struct {
	UserRepository *repositories.UserRepository
	Logger         *zap.Logger
}

func (s *UserServices) CreateUserServices(user models.User) (*models.User, error) {
	// Validation rule here
	if user.FirstName == "" {
		return nil, errors.New("first name cannot be empty")
	}
	if user.LastName == "" {
		return nil, errors.New("last name cannot be empty")
	}
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if user.Address == "" {
		return nil, errors.New("address cannot be empty")
	}

	// Create new user
	newuser := models.User{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Address:     user.Address,
		DateOfBirth: user.DateOfBirth,
	}
	s.Logger.Info("services:= user new created")
	// Add user to the list
	createdUser, err := s.UserRepository.Create(&newuser)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserServices) GetAllUserServices() ([]models.User, error) {
	users, err := s.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
