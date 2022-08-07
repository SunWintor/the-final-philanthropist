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
}

func register(c *gin.Context) {
	var l model.RegisterReq
	if err := c.BindQuery(&l); err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	j := common.NewJWT()
	cl := j.CreateClaims(l.Username + ":" + l.Password)
	token, err := j.CreateToken(cl)
	if err != nil {
		gin_util.FailWithError(c, err)
		return
	}
	common.LoginUser.Put(l.Username, token)
	gin_util.SuccessWithMsg(c, token)
}
