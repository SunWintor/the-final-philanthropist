package model

type UserInfo struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // 密码是md5+盐后的
	Token    string `json:"-"`
}

type UserInfoReply struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (u *UserInfo) ToReply() *UserInfoReply {
	return &UserInfoReply{
		Id:       u.Id,
		Username: u.Username,
		Token:    u.Token,
	}
}

func (u *UserInfo) DeepCopy() *UserInfo {
	if u == nil {
		return nil
	}
	res := &UserInfo{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Token:    u.Token,
	}
	return res
}

type UserIdReq struct {
	UserId int64 `json:"user_id" form:"user_id"`
}
