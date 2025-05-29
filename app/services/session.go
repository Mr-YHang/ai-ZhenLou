package services

import (
	"ai-ZhenLou/app/dao"
	"ai-ZhenLou/app/model"
	"ai-ZhenLou/app/req"
	"ai-ZhenLou/global"
	"context"
)

type Session struct {
	UserDao *dao.User
}

func NewSession(userDao *dao.User) *Session {
	return &Session{
		UserDao: userDao,
	}
}

func (s *Session) Login(ctx context.Context, r *req.LoginReq) (*model.User, error) {
	// 模拟做个查询
	userInfo, err := s.UserDao.FindUserByName(ctx, r.Username)
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Any("user_name", r.Username).Msg("用户登录失败")

		return nil, err
	}

	return userInfo, err
}
