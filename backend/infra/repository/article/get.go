package article

import "backend/domain/model"

func (r *articleRepository) ListArticles() (articles []*model.ResponseArticle, err error) {
	// TODO タグIDをGROUP CONCATで取ると、タグをクリックして、該当タグごとに記事を表示する際にうまく検索できないと思うので、stringのスライスで受け取るようにしたい
	err = r.DB.Table("articles as a").
		Select(`
			a.article_id,
			a.title,
			a.content,
			a.updated_at,
			u.name as user_name,
			GROUP_CONCAT(t.name separator ' / ') as tag
		`).
		Joins("left outer join users as u on u.user_id = a.user_id").
		Joins("left outer join article_tags as at on at.article_id = a.article_id").
		Joins("left outer join tags as t on t.id = at.id").
		Group("a.article_id").
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *articleRepository) ShowArticle(id string) (articles *model.ResponseArticle, err error) {
	err = r.DB.Table("articles as a").
		Select(`
			a.article_id,
			a.title,
			a.content,
			a.updated_at,
			u.name as user_name,
			GROUP_CONCAT(t.name separator ' / ') as tag
		`).
		Joins("left outer join users as u on u.user_id = a.user_id").
		Joins("left outer join article_tags as at on at.article_id = a.article_id").
		Joins("left outer join tags as t on t.id = at.id").
		Where("a.article_id = ?", id).
		Group("a.article_id").
		Find(&articles).Error

	if err != nil {
		return nil, err
	}

	return articles, nil
}
