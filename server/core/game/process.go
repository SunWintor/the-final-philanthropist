package game

import "github.com/SunWintor/tfp/server/core/game/hero"

type Process struct {
	Round        *Round
	PlayerMap    map[string]*Player // 该map初始化后便不会变动，所以无需加锁便可操作。
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
	p.heroSelect()
	p.Round.Next()
}

// todo 选择英雄
func (p *Process) heroSelect() {
	for _, player := range p.PlayerMap {
		h := &hero.Foo2Die{}
		h.Init()
		player.Hero = h
	}
}
