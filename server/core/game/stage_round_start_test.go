package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_roundStart_Run(t *testing.T) {
	ctx := defaultProcessContext()
	StageMap[RoundStartStage].Run(ctx)
	assert.Equal(t, ctx.Round, int64(1))
	assert.Equal(t, ctx.CurrentRoundInfo.RoundNo, int64(1))
	assert.Equal(t, ctx.CurrentRoundInfo.PublicOpinion, int64(PublicOpinionBet+1))
}
