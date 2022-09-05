package service

import (
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) GameEnd(c *gin.Context, game *game.Game) {
	// todo 保存玩家的rank信息

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
