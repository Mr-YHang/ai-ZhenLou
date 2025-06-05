package model

import (
	"github.com/cloudwego/eino/schema"
	"time"
)

type Message struct {
	ID           int64             `json:"id" gorm:"primary_key;column:id"`
	UserID       int64             `json:"user_id" gorm:"column:user_id"`
	DialogueID   string            `json:"dialogue_id" gorm:"column:dialogue_id"`
	DialogueInfo []*schema.Message `json:"dialogue_info" gorm:"serializer:json;column:dialogue_info" `
	CreateAt     time.Time         `json:"create_at" gorm:"column:create_at"`
	UpdateAt     time.Time         `json:"update_at" gorm:"column:update_at"`
}

func (Message) TableName() string {
	return "message"
}
