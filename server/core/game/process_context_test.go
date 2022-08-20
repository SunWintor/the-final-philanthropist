package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_judgementGameEnd(t *testing.T) {
	t.Run("游戏不应该结束", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		ctx.judgementGameEnd()
		assertGaming(t, ctx)
	})
	t.Run("游戏应该结束（只剩一个人）", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		flag := true
		for _, v := range ctx.PlayerMap {
			if flag {
				flag = !flag
				continue
			}
			v.Hero.DecMoney(9999)
		}
		ctx.judgementGameEnd()
		assert.Equal(t, <-ctx.EndGame, struct{}{})
	})
	t.Run("游戏应该结束（全死了）", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		for _, v := range ctx.PlayerMap {
			v.Hero.DecMoney(9999)
		}
		ctx.judgementGameEnd()
		assert.Equal(t, <-ctx.EndGame, struct{}{})
	})
}

func assertGaming(t *testing.T, ctx *ProcessContext) {
	pass := false
	select {
	case <-ctx.EndGame:
		pass = false
	default:
		pass = true
	}
	assert.Equal(t, pass, true)
}

func Test_decPlayerMoney(t *testing.T) {
	t.Run("正常扣减金额", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(11)
		ctx.decPlayerMoney(d, ctx.CurrentRoundInfo.DonatedInfoList[0])
		for _, player := range ctx.PlayerMap {
			if player.PlayerId == ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetMoneyLimit()-d)
			} else {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetCurrentMoney())
			}
		}
	})
}

func Test_decPlayerDonatedMoney(t *testing.T) {
	t.Run("正常同步捐赠金额", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(11)
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = d
		ctx.decPlayerDonatedMoney()
		for _, player := range ctx.PlayerMap {
			if player.PlayerId == ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetMoneyLimit()-d)
			} else {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetCurrentMoney())
			}
		}
	})
}

func Test_decPlayerPunishmentMoney(t *testing.T) {
	t.Run("正常同步舆论惩罚金额", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(11)
		ctx.CurrentRoundInfo.DonatedInfoList[0].PunishmentMoney = d
		ctx.decPlayerPunishmentMoney()
		for _, player := range ctx.PlayerMap {
			if player.PlayerId == ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetMoneyLimit()-d)
			} else {
				assert.Equal(t, player.Hero.GetCurrentMoney(), player.Hero.GetCurrentMoney())
			}
		}
	})
}
