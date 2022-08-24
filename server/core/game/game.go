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
}

const (
	NotStart = 1
	Gaming   = 2
	Ended    = 3
)

func (g *Game) GameInit(playerMap map[string]*Player) {
	g.Status = NotStart
	g.Process = &Process{
		ProcessContext: &ProcessContext{
			Round:     0,
			PlayerMap: playerMap,
			EndGame:   make(chan struct{}),
		},
	}
	g.Process.ProcessContext.initCurrentRound()
}

func (g *Game) Start() <-chan struct{} {
	g.Status = Gaming
	go g.start()
	return g.Process.ProcessContext.EndGame
}

func (g *Game) start() {
	go g.Process.Start()
	<-g.Process.ProcessContext.EndGame
	g.Status = Ended
}

func (g *Game) ToReply(userId int64) *model.GameInfoReply {
	var gameInfo *model.GameInfo
	if g.Process != nil {
		gameInfo = g.Process.ToGameInfoReply(userId)
	}
	return &model.GameInfoReply{
		GameId:   g.GameId,
		RoomId:   g.RoomId,
		Status:   g.Status,
		GameInfo: gameInfo,
	}
}

func (g *Game) Donated(playerId string, donated int64) error {
	if g.Process.Stage.GetStage() != DonatedStage {
		return ecode.StageNotDonatedError
	}
	return g.Process.ProcessContext.CurrentRoundInfo.donated(playerId, donated)
}
