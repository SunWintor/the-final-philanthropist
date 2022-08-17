package model

type GameInfoReq struct {
	UserId int64 `json:"user_id" form:"user_id"`
}

type GameInfoReply struct {
	GameId   string    `json:"game_id"`
	RoomId   string    `json:"room_id"`
	Status   int64     `json:"status"`
	GameInfo *GameInfo `json:"game_info"`
}

type GameInfo struct {
	RoundInfo      *RoundInfo      `json:"round_info"`
	PlayerGameInfo *PlayerGameInfo `json:"player_game_info"`
}

type PlayerGameInfo struct {
	PlayerId         int64           `json:"player_id"`
	Username         string          `json:"username"`
	CurrentMoney     int64           `json:"current_money"`
	HeroInfo         *HeroInfo       `json:"hero_info"`
	RoundHistoryList []*RoundHistory `json:"round_history_list"`
}

type RoundHistory struct {
	RoundNo         int64 `json:"round_no"`
	CurrentMoney    int64 `json:"current_money"`
	DonatedMoney    int64 `json:"donated_money"`
	PunishmentMoney int64 `json:"punishment_money"`
	Bankrupt        bool  `json:"bankrupt"`
}

type HeroInfo struct {
	MoneyLimit int64  `json:"money_limit"`
	Name       string `json:"name"`
	SkillInfo  string `json:"skill_info"`
}

type RoundInfo struct {
	Round         int64  `json:"round"`
	PublicOpinion int64  `json:"public_opinion"`
	Stage         *Stage `json:"stage"`
}

type Stage struct {
	Stage     int64  `json:"stage"`
	StartTime int64  `json:"start_time"`
	Name      string `json:"name"`
	Duration  int64  `json:"duration"`
}
