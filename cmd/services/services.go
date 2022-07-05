package services

import "github.com/iamgoangle/gogolf-template/cmd/domain"

//go:generate mockgen -source=services.go -destination=./mock/repository.go UserRepository

// UserRepository represent database, model, layer interface
type UserRepository interface {
	// Create creates a new user
	Create(usr *domain.UserModel) (*domain.User, error)
}
