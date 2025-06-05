package services

import (
	"ai-ZhenLou/app/dao"
	"ai-ZhenLou/app/model"
	"ai-ZhenLou/global"
	"context"
)

type User struct {
	UserDao *dao.User
}

func NewUser(userDao *dao.User) *User {
	return &User{
		UserDao: userDao,
	}
}

func (s *User) GetUser(ctx context.Context, ID int64) (*model.User, error) {
	userInfo, err := s.UserDao.FindUserByID(ctx, ID)
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Any("user_id", ID).Msg("用户查询失败")

		return nil, err
	}

	return userInfo, err
}
