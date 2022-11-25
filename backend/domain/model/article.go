package model

type Article struct {
	ArticleId string
	UserId    string
	Title     string
	Content   string
}

type Articles []Article
