package model

type JoinRandomRoomReq struct {
	UserID int64 `json:"user_id" form:"user_id"`
}
