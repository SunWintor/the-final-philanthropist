package service

import (
	"github.com/SunWintor/tfp/server/model"
	"testing"
)

func TestService_Login(t *testing.T) {
	t.Run("测试账号不存在的情况下登录", func(t *testing.T) {
		userInfo, err := testSvr.Login(nil, &model.LoginReq{
			Username: "un10",
			Password: "pwd10",
		})
		if userInfo == nil {
			return
		}
		if userInfo.Token != "" || err == nil {
			t.Errorf("没账号也能登录 %+v %v", userInfo, err)
		}
	})
	t.Run("测试正常登录。", func(t *testing.T) {
		userInfo, err := testSvr.Register(nil, &model.RegisterReq{
			Username: "un1",
			Password: "pwd1",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("注册失败。 %+v %v", userInfo, err)
		}
		userInfo, err = testSvr.Login(nil, &model.LoginReq{
			Username: "un1",
			Password: "pwd1",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("没登录成功。 %+v %v", userInfo, err)
		}
	})
	t.Run("测试密码不对情况下登录。", func(t *testing.T) {
		userInfo, err := testSvr.Register(nil, &model.RegisterReq{
			Username: "un20",
			Password: "pwd20",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("注册失败。 %+v %v", userInfo, err)
		}
		userInfo, err = testSvr.Login(nil, &model.LoginReq{
			Username: "un20",
			Password: "pwd100002",
		})
		if userInfo.Token != "" || err == nil {
			t.Errorf("密码错了还登录成功了。 %+v %v", userInfo, err)
		}
	})
}

func TestService_Register(t *testing.T) {
	t.Run("测试初次注册", func(t *testing.T) {
		userInfo, err := testSvr.Register(nil, &model.RegisterReq{
			Username: "un1",
			Password: "pwd1",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("注册失败。 %+v %v", userInfo, err)
		}
	})
	t.Run("测试重复注册", func(t *testing.T) {
		userInfo, err := testSvr.Register(nil, &model.RegisterReq{
			Username: "un2",
			Password: "pwd2",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("注册失败。 %+v %v", userInfo, err)
		}
		userInfo, err = testSvr.Register(nil, &model.RegisterReq{
			Username: "un2",
			Password: "pwd2",
		})
		if userInfo.Token != "" || err == nil {
			t.Errorf("同一个账号重复注册成功。 %+v %v", userInfo, err)
		}
	})
	t.Run("测试账号库不为空的情况下注册", func(t *testing.T) {
		userInfo, err := testSvr.Register(nil, &model.RegisterReq{
			Username: "un3",
			Password: "pwd3",
		})
		if userInfo.Token == "" || err != nil {
			t.Errorf("有账号存在的情况下，注册失败。 %+v %v", userInfo, err)
		}
	})
}
