package dao

import (
	"mars/model"
)

func PrzieRecord(uid uint,prize string) {
	loettery := model.Lottery{
		Uid:uid,
		Prize:prize,
	}
	DB.Create(&loettery)
}

func AddBlackUser(uid uint,room string) error {
	black := model.Blackuser{
		Uid:   uid,
		Room:  room,
	}
	d := DB.Create(&black)
	return d.Error
}

func LoadBlackUsers(room string) (res map[uint]bool) {
	var blacks  []model.Blackuser
	DB.Where("room = ?",room).Find(&blacks)
	res = make(map[uint]bool)
	for _,v := range blacks {
		res[v.Uid] = true
	}
	return
}

func LoadSensitveWords() (res []string) {
	var words  []model.Sensitive_word
	DB.Find(&words )
	for _,v := range words  {
		res = append(res, v.Word)
	}
	return
}
