package service

import (
	"fmt"
	"github.com/SunWintor/tfp/server/core/room"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

// 不返回错误是因为返回了也没有什么用，目前对于该方面的可靠性要求还没有那么高，所以可以这样么干，但要知道这里是有问题的。
func (s *Service) GameEnd(c *gin.Context, r *room.Room) {
	g := r.Game
	for i := 0; i < 5; i++ {
		rankSettlementList, err := s.GameRankSettlement(c, g.Process.ProcessContext.PlayerMap)
		if err != nil {
			log.Printf("GameRankSettlement err %+v %+v", g, err)
			time.Sleep(time.Second * time.Duration(i+1))
			continue
		}
		syncRankUp(rankSettlementList, r)
		break
	}
	return
}

func syncRankUp(rankSettlementList []*rankSettlementTmp, r *room.Room) {
	playerMap := r.Game.Process.ProcessContext.PlayerMap
	for _, v := range rankSettlementList {
		player := playerMap[v.PlayerId]
		player.RankingUp = v.UpdateRanking - v.CurrentRanking
		value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", v.UpdateRanking), 64)
		r.UserMap[v.UserId].Ranking = value
	}
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
