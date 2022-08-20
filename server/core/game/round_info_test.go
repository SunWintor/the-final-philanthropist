package game

import (
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reckonPunishment(t *testing.T) {
	t.Run("只有一个人吃舆论惩罚", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(0)
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = d
		ctx.CurrentRoundInfo.reckonPunishment()
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			if donatedInfo.PlayerId == ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId {
				assert.Equal(t, donatedInfo.PunishmentMoney, ctx.CurrentRoundInfo.PublicOpinion)
			} else {
				assert.Equal(t, donatedInfo.PunishmentMoney, int64(0))
			}
		}
	})
	t.Run("所有人都吃", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		ctx.CurrentRoundInfo.reckonPunishment()
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			assert.Equal(t, donatedInfo.PunishmentMoney, ctx.CurrentRoundInfo.PublicOpinion)
		}
	})
}

func Test_playerByDonated(t *testing.T) {
	t.Run("正常获取", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(99)
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = d
		players := ctx.CurrentRoundInfo.playerByDonated(d)
		assert.Equal(t, players[0], ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId)
	})
	t.Run("获取多个", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		var basePlayers []string
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			basePlayers = append(basePlayers, donatedInfo.PlayerId)
		}

		players := ctx.CurrentRoundInfo.playerByDonated(ctx.CurrentRoundInfo.defaultDonatedMoney())
		assert.Equal(t, players, basePlayers)
	})
	t.Run("一个都没有", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		players := ctx.CurrentRoundInfo.playerByDonated(65535)
		assert.Equal(t, players, []string{})
	})
}

func Test_maxDonated(t *testing.T) {
	t.Run("正常获取", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(99)
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = d
		min := ctx.CurrentRoundInfo.maxDonated()
		assert.Equal(t, min, d)
	})
	t.Run("正常获取，有多个一样的", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		min := ctx.CurrentRoundInfo.maxDonated()
		assert.Equal(t, min, ctx.CurrentRoundInfo.defaultDonatedMoney())
	})
}

func Test_minDonated(t *testing.T) {
	t.Run("正常获取", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(1)
		ctx.CurrentRoundInfo.DonatedInfoList[0].DonatedMoney = d
		min := ctx.CurrentRoundInfo.minDonated()
		assert.Equal(t, min, d)
	})
	t.Run("正常获取，有多个一样的", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		min := ctx.CurrentRoundInfo.minDonated()
		assert.Equal(t, min, ctx.CurrentRoundInfo.defaultDonatedMoney())
	})
}

func Test_donated(t *testing.T) {
	t.Run("正常捐赠", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId

		d := int64(1)
		err := ctx.CurrentRoundInfo.donated(playerId, d)
		assert.Equal(t, err, nil)
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			if donatedInfo.PlayerId == playerId {
				assert.Equal(t, donatedInfo.DonatedMoney, d)
			} else {
				assert.Equal(t, donatedInfo.DonatedMoney, ctx.CurrentRoundInfo.defaultDonatedMoney())
			}
		}
		return
	})
	t.Run("大于最大金钱的捐赠", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)
		playerId := ctx.CurrentRoundInfo.DonatedInfoList[0].PlayerId

		d := int64(9999)
		err := ctx.CurrentRoundInfo.donated(playerId, d)
		assert.Equal(t, err, nil)
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			if donatedInfo.PlayerId == playerId {
				assert.Equal(t, donatedInfo.DonatedMoney, int64(108))
			} else {
				assert.Equal(t, donatedInfo.DonatedMoney, ctx.CurrentRoundInfo.defaultDonatedMoney())
			}
		}
		return
	})
	t.Run("给一个不存在的人捐赠", func(t *testing.T) {
		ctx := defaultProcessContext()
		StageMap[RoundStartStage].Run(ctx)

		d := int64(11)
		err := ctx.CurrentRoundInfo.donated("playerId", d)
		assert.Equal(t, err, ecode.PlayerNotExistsError)
		for _, donatedInfo := range ctx.CurrentRoundInfo.DonatedInfoList {
			assert.Equal(t, donatedInfo.DonatedMoney, ctx.CurrentRoundInfo.defaultDonatedMoney())
		}
		return
	})
}
