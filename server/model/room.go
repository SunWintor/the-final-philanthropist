package model

type RoomInfoReq struct {
	UserId int64 `json:"user_id" form:"user_id"`
}

type RoomInfoReply struct {
	GameId       string      `json:"game_id"`
	RoomId       string      `json:"room_id"`
	Status       int64       `json:"status"`
	RoomUserList []*RoomUser `json:"room_user_list"`
}

type RoomUser struct {
	PlayerId string `json:"player_id"`
	UserId   int64  `json:"user_id"`
	IsReady  bool   `json:"is_ready"`
}

type Hero struct {
	MoneyLimit   int64  `json:"money_limit"`
	CurrentMoney int64  `json:"current_money"`
	Name         string `json:"name"`
	SkillInfo    string `json:"skill_info"`
}
