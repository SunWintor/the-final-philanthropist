package ecode

var (
	AlreadyInRoomError   = newErr(-10000, "该玩家已加入某个房间", "already in a room")
	PlayerNotInRoomError = newErr(-10001, "该玩家不在本房间", "player not in this room")
)
