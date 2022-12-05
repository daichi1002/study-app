package article

import (
	"backend/domain/model"
)

func (r *articleRepository) CreateArticle(article *model.Article) error {
	err := r.DB.Create(&article).Error

	if err != nil {
		return err
	}

	return nil
}
