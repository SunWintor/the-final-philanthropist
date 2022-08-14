package game

import (
	"github.com/SunWintor/tfp/server/core/game/stage"
)

type Game struct {
	GameId  string
	RoomId  string
	Status  int64
	Process *Process
	End     chan struct{}
}

const (
	Gaming = 1
	Ended  = 2
)

func (g *Game) RoundInit(playerMap map[string]*Player) {
	g.Process = &Process{
		Round: &Round{
			CurrentStage: new(stage.Stage), // todo
		},
		PlayerMap: playerMap,
	}
}

func (g *Game) Start() <-chan struct{} {
	g.End = make(chan struct{})
	g.Status = Gaming
	go g.start()
	return g.End
}

func (g *Game) start() {
	go g.Process.Start(g.End)
	<-g.End
	g.Status = Ended
}
