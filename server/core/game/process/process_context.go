package process

import (
	"github.com/SunWintor/tfp/server/core/game"
)

type ProcessContext struct {
	Round               int64
	PlayerMap           map[string]*game.Player // 该map初始化后便不会变动，所以无需加锁便可操作。
	CurrentRoundHistory *RoundHistory
	RoundHistory        []*RoundHistory
	EndGame             <-chan struct{}
}

type RoundHistory struct {
	RoundNo        int64
	PublicOpinion  int64
	DonatedHistory []*DonatedHistory
}

type DonatedHistory struct {
	PlayerId        string
	DonatedMoney    int64
	PunishmentMoney int64
}
