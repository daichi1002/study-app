package repository

import "backend/domain/model"

type UserRepository interface {
	ListUsers() ([]*model.User, error)
}
