package service

import (
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// 不返回错误是因为返回了也没有什么用，目前对于该方面的可靠性要求还没有那么高，所以可以这样么干，但要知道这里是有问题的。
func (s *Service) GameEnd(c *gin.Context, g *game.Game) {
	// todo 保存玩家的rank信息
	for i := 0; i < 5; i++ {
		if err := s.GameRankSettlement(c, g.Process.ProcessContext.PlayerMap); err != nil {
			log.Printf("GameRankSettlement err %+v %+v", g, err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		break
	}
	return
}

func (s *Service) GameInfo(c *gin.Context, arg *model.GameInfoReq) (res *model.GameInfoReply, err error) {
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if err != nil {
		return
	}
	if r.Game != nil {
		res = r.Game.ToReply(arg.UserId)
	}
	return
}

func (s *Service) Donated(c *gin.Context, arg *model.DonatedReq) (err error) {
	var r *room.Room
	r, err = room.UserRoom(arg.UserId)
	if err != nil {
		return
	}
	if r.Status != room.Gaming {
		err = ecode.GameNotStartError
		return
	}
	if arg.Donated < 0 {
		err = ecode.DonatedNotValidError
		return
	}
	if r.Game.Process.ProcessContext.Round != arg.RoundNo {
		err = ecode.RoundAlreadyEndError
		return
	}
	err = r.Game.Donated(arg.PlayerId, arg.Donated)
	return
}
