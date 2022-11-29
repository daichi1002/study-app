package article

import "backend/domain/model"

func (r *articleRepository) ListArticles() (articles []*model.Article, err error) {
	err = r.DB.Table("articles as a").
		Select(`
			a.article_id,
			a.title,
			a.content,
			a.updated_at,
			u.name as user_name,
			mt.name as tag
		`).
		Joins("left outer join users as u on u.user_id = a.user_id").
		Joins("left outer join tags as t on t.article_id = a.article_id").
		Joins("left outer join mst_tag as mt on mt.mst_tag_id = t.tag_id").
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
