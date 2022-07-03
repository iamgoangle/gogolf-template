package services

//go:generate mockgen -source=services.go -destination=./mock/repository.go UserRepository

// UserRepository represent database, model, layer interface
type UserRepository interface {
	// Create creates a new user
	Create() error

	// Update updates by a given user
	Update() error

	// Delete deletes by a given user
	Delete() error
}
