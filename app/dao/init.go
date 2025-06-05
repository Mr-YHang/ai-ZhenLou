package dao

import "ai-ZhenLou/global"

type Dao struct {
	User    *User
	Message *Message
}

func NewDao() *Dao {
	return &Dao{
		User:    NewUser(global.App.DB, global.App.Redis),
		Message: NewMessage(global.App.DB, global.App.Redis),
	}
}
