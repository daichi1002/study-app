package model

import (
	"time"
)

type Article struct {
	Id        string    `json:"articleId" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(50)"`
	Content   string    `json:"content"`
	Tags      []*Tag    `json:"tags" gorm:"many2many:article_tags;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
