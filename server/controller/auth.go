package controller

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func handleAuth(r *gin.Engine) {

	auth := r.Group("/auth")
	auth.POST("/register", register)
	auth.POST("/login", login)
}

func login(c *gin.Context) {
	var l *model.LoginReq
	if err := c.BindQuery(&l); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	l.Password = common.EncryptionPassword(l.Password)
	userInfo, err := svr.Login(c, l)
	if err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	gin_util.SuccessWithObject(c, userInfo)
}

func register(c *gin.Context) {
	var r *model.RegisterReq
	if err := c.ShouldBindJSON(&r); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	r.Password = common.EncryptionPassword(r.Password)
	userInfo, err := svr.Register(c, r)
	if err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	gin_util.SuccessWithObject(c, userInfo)
}
