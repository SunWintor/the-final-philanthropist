package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_donated_end_Run(t *testing.T) {
	t.Run("有人捐赠，有人没动弹。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = 3
		StageMap[DonatedEndStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			if v.PlayerId == playerId {
				assert.Equal(t, v.Hero.GetCurrentMoney(), v.Hero.GetMoneyLimit()-3)
			} else {
				assert.Equal(t, v.Hero.GetCurrentMoney(), v.Hero.GetMoneyLimit()-5)
			}
		}
	})
	t.Run("有人把自己捐赠死了。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = 108
		StageMap[DonatedEndStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			if v.PlayerId == playerId {
				assert.Equal(t, v.Hero.GetCurrentMoney(), v.Hero.GetMoneyLimit()-108)
				assert.Equal(t, v.Hero.IsBankrupt(), true)
			} else {
				assert.Equal(t, v.Hero.GetCurrentMoney(), v.Hero.GetMoneyLimit()-5)
			}
		}
	})
	t.Run("所有人都死了，游戏结束。", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			donatedInfo.DonatedMoney = 108
		}
		StageMap[DonatedEndStage].Run(ctx)
		for _, v := range ctx.PlayerMap {
			assert.Equal(t, v.Hero.GetCurrentMoney(), v.Hero.GetMoneyLimit()-108)
			assert.Equal(t, v.Hero.IsBankrupt(), true)
		}
		assert.Equal(t, <-ctx.EndGame, struct{}{})
	})
}
