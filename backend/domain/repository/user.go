package repository

import "backend/domain/model"

type UserRepository interface {
	ListUsers() ([]*model.User, error)
	GetLoginUser(email string) (*model.User, error)
}
