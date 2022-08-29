package model

type DonatedReq struct {
	UserId   int64  `json:"user_id" form:"user_id"`
	PlayerId string `json:"player_id" form:"player_id"`
	Donated  int64  `json:"donated" form:"donated"`
	RoundNo  int64  `json:"round_no" form:"round_no"`
}

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
	RoundInfo        *RoundInfo          `json:"round_info"`
	PlayerGameInfo   *PlayerGameInfo     `json:"player_game_info"`
	RoundHistoryList []*RoundDonatedInfo `json:"round_history_list"`
}

type PlayerGameInfo struct {
	PlayerId string `json:"player_id"`
	Username string `json:"username"`
	HeroInfo *Hero  `json:"hero_info"`
}

type RoundDonatedInfo struct {
	PlayerId            string  `json:"player_id"`
	Username            string  `json:"username"`
	HeroName            string  `json:"hero_name"`
	CurrentMoneyList    []int64 `json:"current_money_list"`
	DonatedMoneyList    []int64 `json:"donated_money_list"`
	PunishmentMoneyList []int64 `json:"punishment_money_list"`
	BankruptList        []bool  `json:"bankrupt_list"`
}

type RoundInfo struct {
	RoundNo        int64         `json:"round_no"`
	PublicOpinion  int64         `json:"public_opinion"`
	Stage          *Stage        `json:"stage"`
	PlayerInfoList []*PlayerInfo `json:"player_info_list"`
}

type PlayerInfo struct {
	PlayerId        string `json:"player_id"`
	Username        string `json:"username"`
	HeroName        string `json:"hero_name"`
	CurrentMoney    int64  `json:"current_money"`
	DonatedMoney    int64  `json:"donated_money"`
	PunishmentMoney int64  `json:"punishment_money"`
	Bankrupt        bool   `json:"bankrupt"`
}

type Stage struct {
	Stage          int64  `json:"stage"`
	StartTime      int64  `json:"start_time"`
	Name           string `json:"name"`
	DurationSecond int64  `json:"duration"`
}
