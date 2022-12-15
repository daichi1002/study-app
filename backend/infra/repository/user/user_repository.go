package user

import (
	"backend/domain/repository"
	"backend/infra"
)

type userRepository struct {
	*infra.GormHandler
}

// NewArticleRepository NewArticleRepositoryを生成します。.
func NewUserRepository(gormHandler *infra.GormHandler) repository.UserRepository {
	return &userRepository{
		gormHandler,
	}
}
