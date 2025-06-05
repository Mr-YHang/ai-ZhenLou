package handler

import "ai-ZhenLou/app/services"

type Handler struct {
	Session *Session
	AI      *AI
}

func NewHandler(svc *services.Services) *Handler {
	return &Handler{
		Session: NewSession(svc.Session),
		AI:      NewAI(svc.Message, svc.User, svc.Prompt, svc.ChatModel, svc.Tool),
	}
}
