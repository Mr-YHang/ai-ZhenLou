package dao

import (
	"ai-ZhenLou/app/model"
	"context"
	"github.com/cloudwego/eino/schema"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func NewMessage(dB *gorm.DB, redis *redis.Client) *Message {
	return &Message{DB: dB, Redis: redis}
}

func (d *Message) FindMessageHistory(ctx context.Context, userID int64, dialogueID string) (*model.Message, error) {
	res := &model.Message{}

	if err := d.DB.Where("user_id = ? and dialogue_id = ? ", userID, dialogueID).First(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func (d *Message) AddMessage(ctx context.Context, info *model.Message) error {
	return d.DB.Create(info).Error
}

func (d *Message) UpdMessageByUIDAndDID(ctx context.Context, userID int64, dialogueID string, message []*schema.Message) error {
	return d.DB.Where("user_id = ? and dialogue_id = ? ", userID, dialogueID).Updates(map[string]interface{}{
		"update_at":     time.Now(),
		"dialogue_info": message,
	}).Error
}
