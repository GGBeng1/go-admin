package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"userName" gorm:"not null;unique"`
	Password string `json:"password"`
	Nickname string `json:"nickname" gorm:"default:'ggbeng'"`
}
type ResUser struct {
	Username string
	Nickname string
}
