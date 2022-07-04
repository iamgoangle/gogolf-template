package repository

import "github.com/iamgoangle/gogolf-template/cmd/domain"

type MongoDBRepo struct {
}

func NewMongoDB() (*MongoDBRepo, error) {
	return &MongoDBRepo{}, nil
}

func (r *MongoDBRepo) Create(usr *domain.UserModel) (*domain.User, error) {
	return &domain.User{
		Name:  "test",
		Email: "test@test.com",
	}, nil
}
