package model

type LoginReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type RegisterReq struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
