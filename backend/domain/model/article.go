package model

import (
	"backend/constant"
	"time"
)

type Article struct {
	ArticleId string
	UserId    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ResponseArticle struct {
	ArticleId   string `json:"articleId"`
	UserName    string `json:"userName"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Tag         string `json:"tag"`
	UpdatedAt   time.Time
	UpdatedDate string `json:"updatedDate"`
}

type RequestArticle struct {
	ArticleId string
	Title     string       `json:"title"`
	Content   string       `json:"content"`
	Tags      []ArticleTag `json:"tags"`
	UserId    string       `json:"userId"`
}

func (a *ResponseArticle) SetUpdatedDate() {
	a.UpdatedDate = a.UpdatedAt.Format(constant.YYYY_MM_DD)
}
