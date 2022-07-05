package repository

import "github.com/iamgoangle/gogolf-template/cmd/domain"

// MongoDBRepo represent mongodb type
type MongoDBRepo struct{}

type MongoConfigs struct{}

// NewMongoDB instant new mongo repo
// TODO: dependencies injection mongo client
func NewMongoDB(cfg *MongoConfigs) (*MongoDBRepo, error) {
	// mongo logic connection here
	return &MongoDBRepo{}, nil
}

func (r *MongoDBRepo) Create(usr *domain.UserModel) (*domain.User, error) {
	return &domain.User{
		Name:  "test",
		Email: "test@test.com",
	}, nil
}
