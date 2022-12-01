package article

import (
	"backend/domain/model"
	"fmt"
)

func (r *articleRepository) CreateArticle(article *model.Article, tags []*model.ArticleTag) error {

	for _, a := range article.ArticleTags {
		fmt.Println(a)
	}
	err := r.DB.Create(article).Error

	if err != nil {
		return err
	}

	return nil
}
