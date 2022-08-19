package game

import (
	"testing"
)

func Test_gameStart_Run(t *testing.T) {
	ctx := defaultNotStartProcessContext()
	StageMap[GameStartStage].Run(ctx)
	// todo 目前还没有选择英雄
}
