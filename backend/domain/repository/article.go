package repository

import "backend/domain/model"

type ArticleRepository interface {
	ListArticles() (*model.Articles, error)
	CreateArticle(*model.Article) error
}
