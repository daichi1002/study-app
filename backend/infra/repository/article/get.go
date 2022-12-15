package article

import (
	"backend/domain/model"
)

func (r *articleRepository) ListArticles() (articles []*model.Article, err error) {
	err = r.DB.Preload("Tags").Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepository) ShowArticle(id string) (articles *model.Article, err error) {
	err = r.DB.Where("id = ?", id).Preload("Tags").Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
