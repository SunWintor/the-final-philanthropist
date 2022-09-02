package dao

import "database/sql"

type Dao struct {
	db *sql.DB
}

func New() *Dao {
	dao := new(Dao)
	dao.InitMysql()
	return dao
}
