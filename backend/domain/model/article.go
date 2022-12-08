package model

import (
	"time"
)

type Article struct {
	Id        string    `json:"articleId" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"type:varchar(50)" binding:"required,max=30"`
	Content   string    `json:"content" binding:"required"`
	Tags      []*Tag    `json:"tags" gorm:"many2many:article_tags;save_associations:false"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
