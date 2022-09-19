package service

import (
	"database/sql"
	"github.com/SunWintor/tfp/server/core/game"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
)

type rankSettlementTmp struct {
	PlayerId       string
	UserId         int64
	RoomRank       int64
	CurrentRanking float64
	UpdateRanking  float64
	NeedInsert     bool
}

// 天梯段位分算法：
//
// 1.初始所有用户段位分为T=1000，下限1000，赛季末（一个赛季3个月）重置段位分
// 2.每局游戏段位分结算规则：
// (1)计算本局游戏玩家段位分平均值
//     A=(A1+A2+A3+...+AN)/N
//     其中N为本局游戏实际参与人数
// (2)计算得到本局游戏的段位分产出值（随着段位分平均值A的增加，趋向为零和博弈）
//     D = 20000 / A
// (3)计算得到本局游戏的顺位得分初值Sn =200 * ((N-2n+1) / (N-1))
//     其中n为本局游戏特定玩家的顺次排名
// (4)计算得到本局游戏第n名玩家的顺位得分矫正系数 Xn=1（待定）
//     Tn为第n名玩家的初始天梯分
// (5)计算得到本局游戏后玩家的天梯分
//     Tn*=Tn + D/N + Xn * Sn
func updateRanking(rankSettlementList []*rankSettlementTmp) {
	sum := float64(0)
	for _, v := range rankSettlementList {
		sum += v.CurrentRanking
	}
	N := float64(len(rankSettlementList))
	A := sum / N
	D := 20000 / A
	for _, v := range rankSettlementList {
		Xn := float64(1)
		n := float64(v.RoomRank)
		Sn := 200 * ((N - 2*n + 1) / (N - 1))
		Tn := v.CurrentRanking
		v.UpdateRanking = Tn + D/N + Xn*Sn
		log.Printf("N=%+v A=%+v D=%+v Xn=%+v n=%+v Sn=%+v Tn=%+v EndTn=%+v", N, A, D, Xn, n, Sn, Tn, v.UpdateRanking)
		if v.UpdateRanking < model.DefaultRanking {
			v.UpdateRanking = model.DefaultRanking
		}
	}
}

func (s *Service) GameRankSettlement(c *gin.Context, playerMap map[string]*game.Player) (rankSettlementList []*rankSettlementTmp, err error) {
	userIdList := getUserIdList(playerMap)
	var dbRankList []*model.TfpUserRank
	if dbRankList, err = s.dao.RankByUserIdList(c, userIdList); err != nil {
		return
	}
	rankSettlementList = initRankSettlement(dbRankList, playerMap)
	updateRanking(rankSettlementList)
	err = s.dao.TxStartFunc(c, nil, func(tx *sql.Tx) error {
		var err1 error
		for _, v := range rankSettlementList {
			if v.NeedInsert {
				_, err1 = s.dao.CreateUserRank(tx, v.UserId, v.UpdateRanking)
			} else {
				_, err1 = s.dao.UpdateUserRank(tx, v.UserId, v.UpdateRanking)
			}
			if err1 != nil {
				return err1
			}
		}
		return nil
	})
	return
}

func initRankSettlement(dbRankList []*model.TfpUserRank, playerMap map[string]*game.Player) (rankSettlementList []*rankSettlementTmp) {
	rankSettlementList = getRankSettlementList(playerMap)
	rankMap := rankSettlementMap(rankSettlementList)
	for _, v := range dbRankList {
		r := rankMap[v.UserId]
		r.CurrentRanking = v.Ranking
		r.NeedInsert = false
	}
	return
}

func rankSettlementMap(rankList []*rankSettlementTmp) map[int64]*rankSettlementTmp {
	rankMap := make(map[int64]*rankSettlementTmp)
	for _, v := range rankList {
		rankMap[v.UserId] = v
	}
	return rankMap
}

func getUserIdList(playerMap map[string]*game.Player) (userIdList []int64) {
	for _, v := range playerMap {
		userIdList = append(userIdList, v.UserId)
	}
	return
}

func getRankSettlementList(playerMap map[string]*game.Player) (rankSettlementList []*rankSettlementTmp) {
	for _, v := range playerMap {
		rankSettlementList = append(rankSettlementList, &rankSettlementTmp{
			PlayerId:       v.PlayerId,
			UserId:         v.UserId,
			RoomRank:       v.RoomRank,
			CurrentRanking: model.DefaultRanking,
			NeedInsert:     true,
		})
	}
	return
}
