package model

import (
	"math"
	"time"
)

type User struct {
	Id          string       `json:"userId" gorm:"primaryKey;type:varchar(26);size:26"`
	UserName    string       `json:"userName" gorm:"type:varchar(20)"`
	Email       string       `json:"email" gorm:"type:varchar(100)"`
	Password    string       `json:"password" gorm:"type:varchar(100)"`
	BirthDay    time.Time    `json:"birthday" gorm:"type:date"`
	Affiliation string       `json:"affiliation" gorm:"type:varchar(20)"`
	Articles    []*Article   `json:"articles" gorm:"foreignKey:UserId"`
	Languages   []*Languages `json:"languages" gorm:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Languages struct {
	Name  string  `json:"name" gorm:"-"`
	Ratio float64 `json:"ratio" gorm:"-"`
}

func (u *User) SetLanguage(languages []map[string]int) {
	totalLanguage := make(map[string]float64)
	var totalCode float64
	for _, language := range languages {
		if language == nil {
			return
		}
		for key, value := range language {
			if _, ok := totalLanguage[key]; ok {
				totalLanguage[key] = totalLanguage[key] + float64(value)
			} else {
				totalLanguage[key] = float64(value)
			}
			totalCode += float64(value)
		}
	}

	// 桁数を指定して四捨五入
	const baseNumber = 100
	for key, value := range totalLanguage {
		calcNum := value / totalCode * 100
		totalLanguage[key] = math.Round(calcNum*baseNumber) / baseNumber
	}

	for key, value := range totalLanguage {
		val := &Languages{}
		val.Name = key
		val.Ratio = value
		u.Languages = append(u.Languages, val)
	}
}
