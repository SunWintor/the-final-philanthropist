package game

import (
	"github.com/SunWintor/tfp/server/core/game/process"
	"github.com/SunWintor/tfp/server/model"
)

type Game struct {
	GameId  string
	RoomId  string
	Status  int64
	Process *process.Process
	End     chan struct{}
}

const (
	Gaming = 1
	Ended  = 2
)

func (g *Game) ToReply() *model.GameInfoReply {
	return &model.GameInfoReply{
		GameId:   g.GameId,
		RoomId:   g.RoomId,
		Status:   g.Status,
		GameInfo: g.Process.ToGameInfoReply(),
	}
}

func (g *Game) RoundInit(playerMap map[string]*Player) {
	g.Process = &process.Process{
		ProcessContext: &process.ProcessContext{
			Round:     0,
			PlayerMap: playerMap,
			EndGame:   make(chan struct{}),
		},
	}
}

func (g *Game) Start() <-chan struct{} {
	g.End = make(chan struct{})
	g.Status = Gaming
	go g.start()
	return g.End
}

func (g *Game) start() {
	go g.Process.Start()
	<-g.End
	g.Status = Ended
}
