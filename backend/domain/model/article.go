package model

import "time"

type Article struct {
	ArticleId string
	UserId    string
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Articles []Article
