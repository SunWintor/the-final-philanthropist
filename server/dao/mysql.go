package dao

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
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

func (d *Dao) TxStartFunc(ctx *gin.Context, tx *sql.Tx, f func(tx *sql.Tx) error) (err error) {
	if tx == nil {
		tx, err = d.db.BeginTx(context.Background(), nil)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	return f(tx)
}
