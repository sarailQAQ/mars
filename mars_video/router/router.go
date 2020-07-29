package router

import (
	"github.com/gin-gonic/gin"
	"mars/middleware"
)

func SetRouter() {
	r := gin.Default()

	r.POST("/login",Login)
	r.POST("/register",Register)

	v1 := r.Group("/live")
	{
		v1.GET("/:room",middleware.LoginStatus,Enter)
		v1.POST("/start",middleware.LoginStatus,Start)
		v1.POST("/prize",SetPrize)
		v1.POST("/black_list",AddBlackUser)
	}

	r.Run(":8080")
}
