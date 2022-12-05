package model

type ArticleTag struct {
	ArticleId string `gorm:"primaryKey"`
	TagId     int    `gorm:"type:int(11)"`
}

func (a *ArticleTag) SetArticleTag(request *Article) (tags []*ArticleTag) {
	id := request.Id

	for _, tag := range request.Tags {
		newTag := &ArticleTag{ArticleId: id, TagId: tag.Id}
		tags = append(tags, newTag)
	}

	return
}
