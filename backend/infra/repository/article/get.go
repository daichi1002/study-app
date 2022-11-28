package article

import "backend/domain/model"

func (r *articleRepository) ListArticles() (*model.Articles, error) {
	article := &model.Articles{}
	err := r.gormHandler.DB.Find(&article).Error

	if err != nil {
		return nil, err
	}

	return article, nil
}
