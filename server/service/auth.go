package service

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/controller/filter"
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

// Register fixme 当前版本，如果并发去注册同一个账号，会出问题。
// 暂时先不管，等用户量上来了再搞。
func (s *Service) Register(c *gin.Context, arg *model.RegisterReq) (userInfo *model.UserInfo, err error) {
	userInfo = new(model.UserInfo)
	if _, ok := user.GetCopyByName(arg.Username); ok {
		err = ecode.AccountAlreadyRegError
		return
	}
	j := common.NewJWT()
	cl := j.CreateClaims(arg.Username + ":" + arg.Password)
	var token string
	if token, err = j.CreateToken(cl); err != nil {
		return
	}
	userInfo.Id = common.GetNextUserId()
	userInfo.Username = arg.Username
	userInfo.Password = arg.Password
	userInfo.Token = token
	user.Put(userInfo)
	return
}

func (s *Service) Login(c *gin.Context, arg *model.LoginReq) (res *model.UserInfo, err error) {
	var ok bool
	var userInfo *model.UserInfo
	res = new(model.UserInfo)
	if userInfo, ok = user.GetCopyByName(arg.Username); !ok {
		err = ecode.AccountNotExistsError
		return
	}
	if arg.Password != userInfo.Password {
		err = ecode.PasswordError
		return
	}
	_, err = filter.JWTAuth(userInfo.Token)
	res = userInfo
	return
}
