package services

import (
	"ai-ZhenLou/app/dao"
	"ai-ZhenLou/app/model"
	"context"
	"errors"
	"github.com/cloudwego/eino/schema"
	"gorm.io/gorm"
	"time"
)

type Message struct {
	MessageDao *dao.Message
}

func NewMessage(messageDao *dao.Message) *Message {
	return &Message{
		MessageDao: messageDao,
	}
}

// GetMessageHistory 获取会话历史
func (s *Message) GetMessageHistory(ctx context.Context, userID int64, dialogueID string) ([]*schema.Message, error) {
	if userID == 0 || len(dialogueID) == 0 {
		return nil, nil
	}

	messageInfo, err := s.MessageDao.FindMessageHistory(ctx, userID, dialogueID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if messageInfo == nil {
		return nil, nil
	}

	if len(messageInfo.DialogueInfo) > 0 {
		return messageInfo.DialogueInfo, nil
	}

	return nil, nil
}

func (s *Message) UpdMessageHistory(ctx context.Context, ack string, answer *schema.Message, userID int64, dialogueID string) error {
	// step1. 获取会话历史
	history, err := s.GetMessageHistory(ctx, userID, dialogueID)
	if err != nil {
		return err
	}

	ackMessage := &schema.Message{
		Role:    "user",
		Content: ack,
	}

	// step2. 如果为空，则表示为新会话，插入
	if history == nil {
		info := &model.Message{
			UserID:       userID,
			DialogueID:   dialogueID,
			DialogueInfo: []*schema.Message{ackMessage, answer},
			CreateAt:     time.Now(),
			UpdateAt:     time.Now(),
		}

		return s.MessageDao.AddMessage(ctx, info)
	}
	// step3. 如果有历史会话，则需要整合
	history = append(history, ackMessage, answer)

	return s.MessageDao.UpdMessageByUIDAndDID(ctx, userID, dialogueID, history)
}
