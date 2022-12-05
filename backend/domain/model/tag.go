package model

import "time"

type Tag struct {
	Id        int    `json:"tagId" gorm:"primary_key"`
	Name      string `json:"tagName" gorm:"type:varchar(30)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
