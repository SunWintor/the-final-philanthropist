package dao

import (
	"database/sql"
	"time"

	"github.com/SunWintor/tfp/configs"
	_ "github.com/go-sql-driver/mysql"
)

func (d *Dao) InitMysql() {
	dbConf := configs.GetConf().Mysql
	db, err := sql.Open("mysql", dbConf.Url)
	if err != nil {
		panic(err)
	}
	// 最大连接时长
	db.SetConnMaxLifetime(time.Second * dbConf.MaxLifetime)
	// 最大连接数
	db.SetMaxOpenConns(dbConf.MaxOpenConns)
	// 空闲连接数
	db.SetMaxIdleConns(dbConf.MaxIdleConns)
	d.db = db
}
