package dao

import (
	"github.com/SunWintor/tfp/server/ecode"
	"github.com/SunWintor/tfp/server/model"
	"github.com/gin-gonic/gin"
	"log"
)

func (d *Dao) CreateUser(c *gin.Context, user *model.UserInfo) (id int64, err error) {
	result, err := d.db.Exec(`INSERT INTO tfp_user (username, password, token) VALUES (?, ?, ?)`,
		&user.Username, &user.Password, &user.Token)
	if err != nil {
		log.Printf("CreateUser err %+v %+v", id, err)
		return -1, err
	}
	if result != nil {
		id, _ = result.LastInsertId()
	}
	return
}

func (d *Dao) UpdateToken(c *gin.Context, id int64, token string) (err error) {
	stmt, _ := d.db.Prepare(`UPDATE tfp_user SET token = ? WHERE id = ?`)
	defer stmt.Close()
	rows, err := stmt.Query(token, id)
	defer rows.Close()
	var result int
	rows.Scan(&result)
	if err != nil {
		log.Printf("UpdateToken err %+v %+v", id, err)
		return err
	}
	return
}

func (d *Dao) UserById(c *gin.Context, id int64) (user *model.UserInfo, err error) {
	user = new(model.UserInfo)
	err = d.db.QueryRow("SELECT id, username, password, token FROM tfp_user WHERE id = ?", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.Token)
	if err != nil {
		if ecode.IsSqlNoResultErr(err) {
			return nil, nil
		}
		log.Printf("UserById err %+v %+v", id, err)
		return nil, err
	}
	return user, nil
}

func (d *Dao) UserByName(c *gin.Context, username string) (user *model.UserInfo, err error) {
	user = new(model.UserInfo)
	err = d.db.QueryRow("SELECT id, username, password, token FROM tfp_user WHERE username = ?", username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Token)
	if err != nil {
		if ecode.IsSqlNoResultErr(err) {
			return nil, nil
		}
		log.Printf("UserByName err %+v %+v", username, err)
		return nil, err
	}
	return user, nil
}
