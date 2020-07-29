package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"mars/dao"
	"mars/middleware"
	"mars/model"
	"mars/resps"
	"mars/service"
	"net/http"
	"time"
)

func Enter(c *gin.Context) {
	upGrader := websocket.Upgrader{
		CheckOrigin:func(r *http.Request) bool {
			return true
		},
		Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
	}

	conn,err := upGrader.Upgrade(c.Writer, c.Request,nil)
	if err != nil {
		log.Printf("websocket connect error: %s, %s", c.Param("channel"),err.Error())
		return
	}

	var user middleware.UserClaim
	var ok bool
	t,_ := c.Get("user")
	if  user,ok = t.(middleware.UserClaim); !ok{
		resps.Error(c,1002,errors.New("Not login yet"))
		return
	}

	room := c.Param("room")
	client := &service.Client{
		Uid:         user.Id,
		Username:    user.Username,
		Socket:      conn,
		Room:		 room,
		MessageChan: make(chan *service.Message),
	}

	service.WsManager.Register <- client
	go client.Read()
	go client.Write()
	time.Sleep(time.Second*5)
//	resps.Ok(c)
}

func Start(c *gin.Context) {
	var user middleware.UserClaim
	var ok bool
	var roomName string
	t,_ := c.Get("user")
	if  user,ok = t.(middleware.UserClaim); !ok{
		resps.Error(c,1002,errors.New("Not login yet"))
		return
	}
	roomName = c.PostForm("room")
/*	t,_ = c.Get("room")
	if  roomName,ok = t.(string); !ok{
		resps.Error(c,1002,errors.New("room name not found"))
		return
	}*/
	room := service.NewRoom(user.Id,user.Username,roomName)
	service.WsManager.StartRoom <- room
	resps.Ok(c)
}

func SetPrize(c *gin.Context) {
	var prize service.Prize
	if err := c.ShouldBindJSON(&prize); err != nil {
		resps.FormError(c)
		return
	}

	prize.MessageChan = make(chan *service.Message)
	prize.Uids = make(map[uint]string)
	service.WsManager.BeginLottery <- &prize
	go prize.Receive()
	go prize.Timer()
	resps.Ok(c)
}

func AddBlackUser(c *gin.Context) {
	var black model.Blackuser
	if err := c.ShouldBindJSON(&black); err != nil {
		resps.FormError(c)
		return
	}
	err := dao.AddBlackUser(black.Uid,black.Room)

	if err != nil {
		resps.Error(c,1002,err)
		return
	}
	resps.Ok(c)
}