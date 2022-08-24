package game

import (
	"fmt"
	"github.com/SunWintor/tfp/server/common"
)

func defaultProcessContext() *ProcessContext {
	ctx := defaultNotStartProcessContext()
	StageMap[GameStartStage].Run(ctx)
	return ctx
}

func defaultNotStartProcessContext() *ProcessContext {
	playerMap := make(map[string]*Player)
	for i := 0; i < 8; i++ {
		playerId := common.GetRandomPlayerId()
		playerMap[playerId] = &Player{
			UserId:   common.GetNextUserId(),
			PlayerId: playerId,
			Username: fmt.Sprintf("test_user_%d", i),
		}
	}
	ctx := &ProcessContext{
		Round:     0,
		PlayerMap: playerMap,
		EndGame:   make(chan struct{}),
	}
	return ctx
}
