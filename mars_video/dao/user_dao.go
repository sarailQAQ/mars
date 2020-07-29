package dao

import (
	"errors"

	. "mars/model"
)

func Login(username ,password string)uint{
	var u User
	DB.Where(User{
		Username: username,
		Password: password,
	}).First(&u)
	return u.ID
}

func Register(username ,password string)error{
	var u User
	DB.Where("username = ?",username).First(&u)
	if u.ID!=0 {
		return errors.New("username exist!")
	}

	u = User{
		Username: username,
		Password: password,
	}

	DB.Create(&u)
	return nil
}
