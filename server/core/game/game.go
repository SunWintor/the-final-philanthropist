package game

import "github.com/SunWintor/tfp/server/core/game/process"

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

func (g *Game) RoundInit(playerMap map[string]*Player) {
	g.Process = &process.Process{
		Round: &process.Round{
			Stage: process.StageMap[process.GameStartStage],
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
