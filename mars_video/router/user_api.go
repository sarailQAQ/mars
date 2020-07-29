package router

import (
	"errors"
	"github.com/gin-gonic/gin"
	"mars/dao"
	"mars/middleware"
	"mars/resps"
)

type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var f LoginForm
	if err := c.ShouldBindJSON(&f); err != nil {
		resps.FormError(c)
		return
	}

	id :=dao.Login(f.Username,f.Password)

	if id != 0 {
		token := middleware.Creat(f.Username,id)
		resps.OkWithData(c,gin.H{
			"token":token,
			"username": f.Username,
		})	// login success
		return
	}
	resps.Error(c,1001,errors.New("password error"))
}

func Register(c *gin.Context) {
	var f LoginForm
	if err := c.ShouldBindJSON(&f); err != nil {
		resps.FormError(c)
		return
	}

	err := dao.Register(f.Username,f.Password)

	if err == nil {
		resps.OkWithData(c, gin.H{"message":"register successfully"})
		return
	}
	resps.Error(c,1001,errors.New("register failed"))
}
