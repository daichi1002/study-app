package tag

import (
	"backend/domain/repository"
	"backend/infra"
)

type tagRepository struct {
	*infra.GormHandler
}

// NewArticleRepository NewArticleRepositoryを生成します。.
func NewTagRepository(gormHandler *infra.GormHandler) repository.TagRepository {
	return &tagRepository{
		gormHandler,
	}
}
