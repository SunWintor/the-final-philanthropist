package service

import (
	"fmt"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"sort"
)

func (s *Service) GameEnd(c *gin.Context, g *game.Game) {
	// todo 保存玩家的rank信息
	m := g.Process.ProcessContext.PlayerMap
	var playerList []*game.Player
	for _, v := range m {
		playerList = append(playerList, v)
	}
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Hero.GetCurrentMoney() > playerList[j].Hero.GetCurrentMoney()
	})
	for _, v := range playerList {
		fmt.Printf("%+v\n", v.Hero)
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
