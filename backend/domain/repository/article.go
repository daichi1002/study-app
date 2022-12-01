package repository

import (
	"backend/domain/model"
)

type ArticleRepository interface {
	ListArticles() ([]*model.ResponseArticle, error)
	CreateArticle(article *model.Article, tags []*model.ArticleTag) error
}
