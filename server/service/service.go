package service

import "github.com/SunWintor/tfp/server/dao"

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	service := new(Service)
	service.dao = dao.New()
	return service
}
