package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_round_end_Run(t *testing.T) {
	t.Run("回合正常结束", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		StageMap[RoundEndStage].Run(ctx)
		assert.Equal(t, ctx.CurrentRoundInfo, ctx.RoundHistory[0])
		return
	})
	t.Run("游戏结束（全死了）", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			v.Hero.DecMoney(108)
		}
		StageMap[RoundEndStage].Run(ctx)
		assert.Equal(t, ctx.CurrentRoundInfo, ctx.RoundHistory[0])
		assert.Equal(t, <-ctx.EndGame, struct{}{})
		return
	})
	t.Run("游戏结束（还剩一个）", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		flag := true
		for _, v := range ctx.PlayerMap {
			if flag {
				flag = !flag
				continue
			}
			v.Hero.DecMoney(108)
		}
		StageMap[RoundEndStage].Run(ctx)
		assert.Equal(t, ctx.CurrentRoundInfo, ctx.RoundHistory[0])
		assert.Equal(t, <-ctx.EndGame, struct{}{})
		return
	})
	t.Run("游戏不结束（还剩两个）", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		count := 0
		for _, v := range ctx.PlayerMap {
			if count < 2 {
				count++
				continue
			}
			v.Hero.DecMoney(108)
		}
		StageMap[RoundEndStage].Run(ctx)
		assert.Equal(t, ctx.CurrentRoundInfo, ctx.RoundHistory[0])
		assertGaming(t, ctx)
		return
	})
}
