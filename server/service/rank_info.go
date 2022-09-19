package service

import (
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetUserRankInfo(ctx *gin.Context, userId int64) (rankInfo *model.TfpUserRank, err error) {
	if rankInfo, err = s.dao.RankByUserId(ctx, userId); err != nil {
		return
	}
	if rankInfo == nil {
		rankInfo = &model.TfpUserRank{
			UserId:  userId,
			Ranking: model.DefaultRanking,
		}
	}
	return
}
