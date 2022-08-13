package model

type UserInfo struct {
	Id       int64  `json:"Id"`
	Username string `json:"username"`
	Password string `json:"-"` // 密码是md5+盐后的
	Token    string `json:"-"`
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

type UIdReq struct {
	UId int64 `json:"uId" form:"uId"`
}
