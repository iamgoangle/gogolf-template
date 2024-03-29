package services

import (
	"fmt"

	"github.com/iamgoangle/gogolf-template/cmd/domain"
)

// UserService represent user service type
type UserService struct {
	usrRepo UserRepository
}

// NewUserService instant a new user service
func NewUserService(rpUsr UserRepository) (*UserService, error) {
	// service layer logic ....

	return &UserService{
		usrRepo: rpUsr,
	}, nil
}

func (s *UserService) Create(*domain.User) (*domain.User, error) {
	usrRepo, err := s.usrRepo.Create(&domain.UserModel{})
	if err != nil {
		return nil, fmt.Errorf("[UserService.Create]: unable to create new user %w", err)
	}

	return &domain.User{
		Name:  usrRepo.Name,
		Email: usrRepo.Email,
	}, nil
}

func (s *UserService) Delete() error {
	return nil
}
