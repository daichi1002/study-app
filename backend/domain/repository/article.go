package repository

import (
	"backend/domain/model"
)

type ArticleRepository interface {
	ListArticles() ([]*model.Article, error)
	CreateArticle(*model.Article) error
	ShowArticle(string) (*model.Article, error)
	UpdateArticle(*model.Article) error
	DeleteArticle(string) error
}
