package handler

import (
	"ai-ZhenLou/app/req"
	"ai-ZhenLou/app/resp"
	"ai-ZhenLou/app/services"
	"ai-ZhenLou/global"
	"github.com/gin-gonic/gin"
)

type Session struct {
	SessionSvc *services.Session
}

func NewSession(sessionSvc *services.Session) *Session {
	return &Session{
		SessionSvc: sessionSvc,
	}
}

// Login 登录
func (h *Session) Login(c *gin.Context) {
	var (
		r   req.LoginReq
		ctx = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&r); err != nil {
		resp.Fail(c, global.ParamErrCode, global.ParamErrMsg)
		return
	}

	if err := r.Check(); err != nil {
		resp.Fail(c, global.ParamErrCode, err.Error())
		return
	}

	res, err := h.SessionSvc.Login(ctx, &r)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}

	resp.Success(c, res)
}
