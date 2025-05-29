package services

import (
	"ai-ZhenLou/app/req"
	"ai-ZhenLou/app/resp"
	"ai-ZhenLou/global"
	"context"
	"errors"
)

type Session struct{}

func NewSession() *Session {
	return &Session{}
}

func (s *Session) Login(ctx context.Context, r *req.LoginReq) (*resp.LoginResp, error) {
	// 这是一个模拟报错
	err := errors.New("数据库查询失败")

	// 打印日志
	global.App.Log.Err(err).Any("user_name", r.Username).Msg("用户登录失败")

	return nil, err
}
