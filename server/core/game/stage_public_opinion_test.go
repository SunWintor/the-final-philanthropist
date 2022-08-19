package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_publicOpinion_Run(t *testing.T) {
	t.Run("所有人都没操作。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		StageMap[PublicOpinionStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			assert.Equal(t, v.Hero.GetCurrentMoney(), int64(108-3))
		}
	})
	t.Run("只有一个人最低。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = 0
		StageMap[PublicOpinionStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			if v.PlayerId == playerId {
				assert.Equal(t, v.Hero.GetCurrentMoney(), int64(108-3))
			} else {
				assert.Equal(t, v.Hero.GetCurrentMoney(), int64(108))
			}
		}
	})
	t.Run("两个人一样都是最低。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId
		playerId2 := ctx.CurrentRoundInfo.DonatedInfoList[1].PlayerId
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = 0
		ctx.CurrentRoundInfo.DonatedInfoList[1].DonatedMoney = 0
		StageMap[PublicOpinionStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			if v.PlayerId == playerId || v.PlayerId == playerId2 {
				assert.Equal(t, v.Hero.GetCurrentMoney(), int64(108-3))
			} else {
				assert.Equal(t, v.Hero.GetCurrentMoney(), int64(108))
			}
		}
	})
}
