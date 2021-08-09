package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username    string `json:"username"`
	Age         int64  `json:"age"`
	Gender      bool   `json:"gender"` // 1 = Male - 0 = Female
	Nationality string `json:"nationality"`
}
