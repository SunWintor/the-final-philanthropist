package process

import (
	"github.com/SunWintor/tfp/server/core/game"
)

type Process struct {
	Round        *Round
	PlayerMap    map[string]*game.Player // 该map初始化后便不会变动，所以无需加锁便可操作。
	RoundHistory []*RoundHistory
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

func (p *Process) Start(end chan<- struct{}) {
	stageEnd := p.Round.NextStage()
	// todo 针对每种不同的stage，进行对应的操作
}
