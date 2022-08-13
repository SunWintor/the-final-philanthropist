package model

type JoinRandomRoomReq struct {
	UserId int64 `json:"user_Id" form:"user_Id"`
}

type JoinRandomRoomReply struct {
	RoomId string `json:"room_Id" form:"room_Id"`
}
