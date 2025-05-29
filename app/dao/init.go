package dao

import "ai-ZhenLou/global"

type Dao struct {
	User *User
}

func NewDao() *Dao {
	return &Dao{
		User: NewUser(global.App.DB, global.App.Redis),
	}
}
