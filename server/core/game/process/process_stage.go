package process

import "github.com/SunWintor/tfp/server/core/game/hero"

// todo 选择英雄
func (p *Process) OnGameStart(end chan<- struct{}) {
	for _, player := range p.PlayerMap {
		h := &hero.Foo2Die{}
		h.Init()
		player.Hero = h
	}
}
