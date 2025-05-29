package services

import "ai-ZhenLou/app/dao"

type Services struct {
	Session *Session
}

func NewServices(dao *dao.Dao) *Services {
	return &Services{
		Session: NewSession(dao.User),
	}
}
