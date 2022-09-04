package service

import (
	"github.com/SunWintor/tfp/server/core/user"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Service) GetUserByUsername(c *gin.Context, username string) (userInfo *model.UserInfo, err error) {
	var ok bool
	if userInfo, ok = user.GetCopyByName(username); ok {
		return
	}
	if userInfo, err = s.dao.UserByName(c, username); err != nil {
		err = errors.WithStack(err)
		return
	}
	if userInfo != nil && userInfo.Id > 0 {
		user.Put(userInfo)
		return
	}
	return
}

func (s *Service) CreateUser(c *gin.Context, userInfo *model.UserInfo) (err error) {
	var existsUser *model.UserInfo
	if existsUser, err = s.GetUserByUsername(c, userInfo.Username); err != nil {
		err = errors.WithStack(err)
		return
	}
	if existsUser != nil && existsUser.Id > 0 {
		err = ecode.AccountAlreadyRegError
		return
	}
	if userInfo.Id, err = s.dao.CreateUser(c, userInfo); err != nil {
		err = errors.WithStack(err)
		return
	}
	user.Put(userInfo)
	return
}
