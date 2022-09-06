package repository

import "github.com/iamgoangle/gogolf-template/cmd/domain"

type Database interface {
	Create(*domain.UserModel) error
}
