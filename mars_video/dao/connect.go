package dao

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"log"
	"mars/model"
)

var (
	DB *gorm.DB
	Conn redis.Conn
)


func Init() {
	sql, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/mars?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	DB = sql
	if !HasTable(model.User{}){
		panic("can not load table users")
	}
	if !HasTable(model.Blackuser{}){
		panic("can not load table blackusers")
	}
	if !HasTable(model.Lottery{}){
		panic("can not load table lotterys")
	}
	if !HasTable(model.Sensitive_word{}){
		panic("can not load table sensitive_words")
	}
}
func HasTable(data interface{}) bool {
	if !DB.HasTable(data) {
		dd:=DB.CreateTable(data)
		if dd.Error != nil {
			return false
		}
	}
	return true
}

func RedisInit() bool {
	c,err := redis.Dial("tcp","127.0.0.1:6379")
	if err != nil {
		log.Println(err)
		return false
	}
	Conn = c
	return true
}
