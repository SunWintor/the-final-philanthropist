package service

import (
	"fmt"
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/game/hero"
	"testing"
)

func TestService_GameEnd(t *testing.T) {
	t.Run("测试账号不存在的情况下登录", func(t *testing.T) {
		m := make(map[string]*game.Player)
		for i := 0; i < 8; i++ {
			playerId := common.GetRandomPlayerId()
			userId := common.GetNextUserId()
			username := fmt.Sprintf("%s:%d", playerId, userId)
			h := &hero.Foo2Die{}
			h.Init()
			h.CurrentMoney = int64(i - 4)
			m[playerId] = &game.Player{
				UserId:   userId,
				PlayerId: playerId,
				Username: username,
				Hero:     h,
			}
		}
		testSvr.GameEnd(nil, &game.Game{
			Process: &game.Process{
				ProcessContext: &game.ProcessContext{
					PlayerMap: m,
				},
			},
		})
	})
}
