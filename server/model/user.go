package model

type UserInfo struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // 密码是md5+盐后的
	Token    string `json:"-"`
}

func (u *UserInfo) DeepCopy() *UserInfo {
	if u == nil {
		return nil
	}
	res := &UserInfo{
		ID:       u.ID,
		Username: u.Username,
		Password: u.Password,
		Token:    u.Token,
	}
	return res
}

type UIDReq struct {
	UID int64 `json:"uid" form:"uid"`
}
