package model

import "time"

// TfpUserRank 用户排名
type TfpUserRank struct {
	ID      int64     `json:"id"`      // id
	UserId  int64     `json:"user_id"` // 用户id
	Ranking float64   `json:"ranking"` // 排位分
	Ver     int64     `json:"ver"`     // 版本号，用于幂等
	Ctime   time.Time `json:"ctime"`   // 创建时间
	Mtime   time.Time `json:"mtime"`   // 更新时间
}
