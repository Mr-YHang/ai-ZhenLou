package dao

import (
	"ai-ZhenLou/app/model"
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type User struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewUser(dB *gorm.DB, redis *redis.Client) *User {
	return &User{DB: dB, Redis: redis}
}

func (d *User) FindUserByName(ctx context.Context, name string) (*model.User, error) {
	res := &model.User{}

	if err := d.DB.Where("user_name = ?", name).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (d *User) FindUserByID(ctx context.Context, ID int64) (*model.User, error) {
	res := &model.User{}

	if err := d.DB.Where("id = ?", ID).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}
