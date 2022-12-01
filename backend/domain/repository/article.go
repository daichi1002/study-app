package repository

import (
	"backend/domain/model"
)

type ArticleRepository interface {
	ListArticles() ([]*model.ResponseArticle, error)
	CreateArticle(*model.Article) error
	ShowArticle(string) (*model.ResponseArticle, error)
}
