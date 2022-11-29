package repository

import (
	"backend/domain/model"
)

type ArticleRepository interface {
	ListArticles() ([]*model.Article, error)
	CreateArticle(*model.Article) error
}
