package model

type ArticleTag struct {
	ArticleId string
	TagId     int
}

func (a *ArticleTag) SetArticleTag(request *RequestArticle) (tags []*ArticleTag) {
	id := request.ArticleId

	for _, tag := range request.ArticleTags {
		newTag := &ArticleTag{ArticleId: id, TagId: tag}
		tags = append(tags, newTag)
	}

	return
}
