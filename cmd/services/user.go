package services

import "github.com/iamgoangle/gogolf-template/cmd/domain"

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

func (s *UserService) Create() (*domain.User, error) {
	return &domain.User{
		Name:  "test",
		Email: "test@test.com",
	}, nil
}
