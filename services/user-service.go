package services

import (
	"ManagementSystem/persist"
	"ManagementSystem/persist/models"
	"fmt"
	"github.com/google/uuid"
)

type UserService struct {
	dao *persist.UserDao
}

func NewUserService(dao *persist.UserDao) *UserService {
	return &UserService{dao: dao}
}

func (s *UserService) RegisterUser(firstname, lastname, email string) (*models.UserDbo, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("an error creating uuid for new user with name %s %s and email %s occured. Reason was %w", firstname, lastname, email, err)
	}
	user := &models.UserDbo{FirstName: firstname, LastName: lastname, Email: email, ID: id}
	if err := s.dao.CreateUser(user); err != nil {
		return nil, fmt.Errorf("could not register user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetUser(id uuid.UUID) (*models.UserDbo, error) {
	return s.dao.GetUser(id)
}

func (s *UserService) GetAllUsers() ([]models.UserDbo, error) {
	return s.dao.GetAllUsers()
}

func (s *UserService) DeleteUser(id uuid.UUID) error {
	if err := s.dao.DeleteUser(id); err != nil {
		return fmt.Errorf("could not delete user with id %s: %w", id, err)
	}
	return nil
}

func (s *UserService) UpdateUser(id uuid.UUID, firstname, lastname, email string) (*models.UserDbo, error) {
	user, err := s.dao.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve user with id %s: %w", id, err)
	}
	user.FirstName = firstname
	user.LastName = lastname
	user.Email = email
	if err := s.dao.UpdateUser(user); err != nil {
		return nil, fmt.Errorf("could not update user with id %s: %w", id, err)
	}
	return user, nil

}
