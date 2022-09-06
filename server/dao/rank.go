package dao

import (
	"database/sql"
	"log"
)

func (d *Dao) CreateUserRank(tx *sql.Tx, userId, ranking int64) (id int64, err error) {
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
