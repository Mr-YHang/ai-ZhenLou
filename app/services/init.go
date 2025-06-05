package services

import "ai-ZhenLou/app/dao"

type Services struct {
	Session   *Session
	Message   *Message
	User      *User
	Prompt    *Prompt
	Tool      *Tool
	ChatModel *ChatModel
}

func NewServices(dao *dao.Dao) *Services {
	return &Services{
		Session:   NewSession(dao.User),
		Message:   NewMessage(dao.Message),
		User:      NewUser(dao.User),
		Prompt:    NewPrompt(),
		Tool:      NewTool(),
		ChatModel: NewChatModel(),
	}
}
