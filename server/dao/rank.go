package dao

import (
	"database/sql"
	"fmt"
	"github.com/SunWintor/tfp/server/common"
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
)

func (d *Dao) CreateUserRank(tx *sql.Tx, userId int64, ranking float64) (id int64, err error) {
	result, err := d.db.Exec(`INSERT INTO tfp_user_rank (user_id, ranking) VALUES (?, ?)`, &userId, &ranking)
	if err != nil {
		log.Printf("CreateUserRank err %+v %+v", id, err)
		return -1, err
	}
	if result != nil {
		id, _ = result.LastInsertId()
	}
	return
}

func (d *Dao) UpdateUserRank(tx *sql.Tx, userId int64, ranking float64) (id int64, err error) {
	result, err := d.db.Exec(`UPDATE tfp_user_rank SET ranking = ? WHERE user_id = ?`, &ranking, &userId)
	if err != nil {
		log.Printf("UpdateUserRank err %+v %+v", id, err)
		return -1, err
	}
	if result != nil {
		id, _ = result.RowsAffected()
	}
	return
}

func (d *Dao) RankByUserIds(c *gin.Context, userIdList []int64) (rankList []*model.TfpUserRank, err error) {
	rows, err := d.db.Query(fmt.Sprintf("select id, user_id, ranking, ver, ctime, mtime from tfp_user_rank where user_id in(%s)", common.DisorderString(userIdList)))
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var rank *model.TfpUserRank
		err = rows.Scan(&rank.ID, &rank.UserId, &rank.Ranking, &rank.Ver, &rank.Ctime, &rank.Mtime)
		if err != nil && !ecode.IsSqlNoResultErr(err) {
			log.Printf("RankByUserIds err %+v %+v", userIdList, err)
			return nil, err
		}
	}
	return
}
