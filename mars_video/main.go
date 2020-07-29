package main

import (
	"mars/dao"
	"mars/router"
	"mars/service"
)

func main() {
	dao.Init()
//	dao.RedisInit()
	go service.WsManager.Work()
	go service.LManager.Work()
	router.SetRouter()
}
