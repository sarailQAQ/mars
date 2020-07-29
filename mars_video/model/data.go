package model

import (
	"github.com/jinzhu/gorm"
)

//数据库结构体

type Sensitive_word struct {
	Word string `gorm:"primary_key" sql:"index"`
}

type User struct {
	gorm.Model
	Username string
	Password string
}

type Blackuser struct {
	gorm.Model
	Uid uint `json:"uid"`
	Room string `sql:"index" json:"room"`
}

type Lottery struct {
	gorm.Model
	Uid uint
	Prize string
}




