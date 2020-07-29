package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mars/resps"
)

type UserClaim struct {
	Id uint
	Username string
}

func LoginStatus(c *gin.Context) {
	auth:= c.GetHeader("token")

	if len(auth) < 7 {
		resps.Error(c,2,errors.New("Illegal jwt"))
		c.Abort()
		return
	}
	//token := auth[7:]
	uid,user,err := Check(auth)
	if err != nil {
		resps.Error(c,2,err)
		c.Abort()
		return
	}

	c.Set("user",UserClaim{
		Id:       uid,
		Username: user,
	})

	c.Next()
	return
}

/*
//判断是否是房间拥有者
func OpenRoom(c *gin.Context) {
	var user UserClaim
	var ok bool
	var roomName string
	t,_ := c.Get("user")
	if  user,ok = t.(UserClaim); !ok{
		resps.Error(c,1002,errors.New("Not login yet"))
		return
	}
	roomName = c.PostForm("room")


}*/