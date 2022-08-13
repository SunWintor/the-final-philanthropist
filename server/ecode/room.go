package ecode

var (
	AlreadyInRoomError      = newErr(10000, "该玩家已加入某个房间", "already in a room")
	PlayerRoomWrongError    = newErr(10001, "该玩家不在本房间", "player room is not corresponding")
	RoomIsFullError         = newErr(10002, "房间已满", "room is full")
	RoomNotExistsError      = newErr(10003, "房间不存在", "room not exists")
	PlayerNotInRoomError    = newErr(10004, "该玩家未加入任何房间", "not in a room")
	RoomNotReadyError       = newErr(10005, "房间未处于准备阶段", "room not ready")
	InsufficientPlayerError = newErr(10006, "玩家人数不足", "insufficient player")
	PlayerNotReadyError     = newErr(10007, "有玩家尚未准备", "some players are not ready")
)
