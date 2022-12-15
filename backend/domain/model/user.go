package model

import "time"

type User struct {
	Id          string     `json:"userId" gorm:"primaryKey;type:varchar(26);size:26"`
	UserName    string     `json:"userName" gorm:"type:varchar(20)"`
	Email       string     `json:"email" gorm:"type:varchar(100)"`
	Password    string     `json:"password" gorm:"type:varchar(100)"`
	BirthDay    time.Time  `json:"birthday" gorm:"type:date"`
	Affiliation string     `json:"affiliation" gorm:"type:varchar(20)"`
	Articles    []*Article `json:"articles" gorm:"foreignKey:UserId"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
