package model

import (
	"backend/constant"
	"time"
)

type Article struct {
	ArticleId   string `json:"articleId"`
	UserName    string `json:"userName"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Tag         string `json:"tag"`
	UpdatedAt   time.Time
	UpdatedDate string `json:"updatedDate"`
}

func (a *Article) SetUpdatedDate() {
	a.UpdatedDate = a.UpdatedAt.Format(constant.YYYY_MM_DD)
}
