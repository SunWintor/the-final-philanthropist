package model

type RoomInfoReply struct {
	GameId       string         `json:"game_id"`
	RoomId       string         `json:"room_id"`
	Status       int64          `json:"status"`
	RoomUserList []*RoomUser    `json:"room_user_list"`
	GameInfo     *GameInfoReply `json:"game_info"`
}

type RoomUser struct {
	UserId    int64   `json:"user_id"`
	Username  string  `json:"username"`
	IsReady   bool    `json:"is_ready"`
	Ranking   float64 `json:"ranking"`
	RankingUp float64 `json:"ranking_up"`
}

type Hero struct {
	MoneyLimit   int64  `json:"money_limit"`
	CurrentMoney int64  `json:"current_money"`
	Name         string `json:"name"`
	SkillInfo    string `json:"skill_info"`
}

type JoinRoomIdReq struct {
	UserId int64  `json:"user_id" form:"user_id"`
	RoomId string `json:"room_id" form:"room_id"`
}
