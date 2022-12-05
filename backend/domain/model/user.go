package model

import "time"

type User struct {
	UserId    string `gorm:"primaryKey"`
	UserName  string `gorm:"type:varchar(20)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
