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

func (r *articleRepository) UpdateArticle(article *model.Article) error {
	// TODO：タグの変更で元のタグが削除された場合の対応をもう少し考えたい
	deleteErr := r.DB.Debug().Where("article_id = ?", article.Id).Delete(&model.ArticleTag{}).Error

	if deleteErr != nil {
		return deleteErr
	}

	err := r.DB.Debug().Save(&article).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *articleRepository) DeleteArticle(id string) (err error) {
	err = r.DB.Debug().Where("article_id = ?", id).Delete(&model.ArticleTag{}).Error

	if err != nil {
		return err
	}

	err = r.DB.Debug().Where("id = ?", id).Delete(&model.Article{}).Error

	if err != nil {
		return err
	}

	return nil
}
