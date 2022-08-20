package controller

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/controller/gin_util"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
)

func handleAuth(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/register", register)
	auth.POST("/login", login)
}

func login(c *gin.Context) {
	var l *model.LoginReq
	if err := c.ShouldBindJSON(&l); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	log.Printf("%+v", l)
	l.Password = common.EncryptionPassword(l.Password)
	userInfo, err := svr.Login(c, l)
	gin_util.AutoResult(c, userInfo, err)
}

func register(c *gin.Context) {
	var r *model.RegisterReq
	if err := c.ShouldBindJSON(&r); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	r.Password = common.EncryptionPassword(r.Password)
	userInfo, err := svr.Register(c, r)
	gin_util.AutoResult(c, userInfo, err)
}
