package service

import (
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/controller/filter"
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Register fixme 当前版本，如果并发去注册同一个账号，会出问题。
// 暂时先不管，等用户量上来了再搞。
func (s *Service) Register(c *gin.Context, arg *model.RegisterReq) (res *model.UserInfoReply, err error) {
	userInfo := new(model.UserInfo)
	res = new(model.UserInfoReply)
	if userInfo, err = s.GetUserByUsername(c, arg.Username); err != nil {
		return
	}
	if userInfo != nil && userInfo.Id > 0 {
		err = ecode.AccountAlreadyRegError
		return
	}
	if userInfo == nil {
		userInfo = new(model.UserInfo)
	}
	j := common.NewJWT()
	if userInfo.Token, err = j.CreateToken(j.CreateClaims(arg.Username + ":" + arg.Password)); err != nil {
		return
	}
	userInfo.Username = arg.Username
	userInfo.Password = arg.Password
	if err = s.CreateUser(c, userInfo); err != nil {
		err = errors.WithStack(err)
		return
	}
	res = userInfo.ToReply()
	return
}

func (s *Service) Login(c *gin.Context, arg *model.LoginReq) (res *model.UserInfoReply, err error) {
	var userInfo *model.UserInfo
	res = new(model.UserInfoReply)

	if userInfo, err = s.GetUserByUsername(c, arg.Username); err != nil {
		return
	}
	if userInfo == nil || userInfo.Id < 0 {
		err = ecode.AccountNotExistsError
		return
	}
	if arg.Password != userInfo.Password {
		err = ecode.PasswordError
		return
	}
	j := common.NewJWT()
	if userInfo.Token, err = j.CreateToken(j.CreateClaims(arg.Username + ":" + arg.Password)); err != nil {
		return
	}
	if err = s.dao.UpdateToken(c, userInfo.Id, userInfo.Token); err != nil {
		return
	}
	_, err = filter.JWTAuth(userInfo.Token)
	user.Put(userInfo)
	res = userInfo.ToReply()
	return
}
