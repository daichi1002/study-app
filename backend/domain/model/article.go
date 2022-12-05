package model

import (
	"backend/constant"
	"time"
)

type Article struct {
	ArticleId string
	// ユーザーは別で取ってくるから不要になる
	UserId  string
	Title   string
	Content string
	// tagsをもつ
	ArticleTags []*ArticleTag `gorm:"foreignKey:ArticleId;references:ArticleId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// 消す
type ResponseArticle struct {
	ArticleId   string `json:"articleId"`
	UserName    string `json:"userName"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Tag         string `json:"tag"`
	UpdatedAt   time.Time
	UpdatedDate string `json:"updatedDate"`
}

// 消す
type RequestArticle struct {
	ArticleId   string
	UserId      string `json:"userId"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	ArticleTags []int  `json:"tags"`
}

func (a *ResponseArticle) SetUpdatedDate() {
	a.UpdatedDate = a.UpdatedAt.Format(constant.YYYY_MM_DD)
}

func (a *Article) SetArticle(request *RequestArticle) {
	a.ArticleId = request.ArticleId
	a.UserId = request.UserId
	a.Title = request.Title
	a.Content = request.Content

	for _, tag := range request.ArticleTags {
		newTag := &ArticleTag{TagId: tag}
		a.ArticleTags = append(a.ArticleTags, newTag)
	}
}
