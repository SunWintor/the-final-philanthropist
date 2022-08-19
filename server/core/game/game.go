package game

import (
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
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

func (g *Game) GameInit(playerMap map[string]*Player) {
	g.Process = &Process{
		ProcessContext: &ProcessContext{
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

func (g *Game) ToReply(userId int64) *model.GameInfoReply {
	return &model.GameInfoReply{
		GameId:   g.GameId,
		RoomId:   g.RoomId,
		Status:   g.Status,
		GameInfo: g.Process.ToGameInfoReply(userId),
	}
}

func (g *Game) Donated(playerId string, donated int64) error {
	if g.Process.Stage.GetStage() != DonatedStage {
		return ecode.StageNotDonatedError
	}
	return g.Process.ProcessContext.CurrentRoundInfo.Donated(playerId, donated)
}
