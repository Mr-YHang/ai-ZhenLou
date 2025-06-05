package services

import (
	"ai-ZhenLou/app/common"
	"ai-ZhenLou/global"
	"context"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type Prompt struct {
}

func NewPrompt() *Prompt {
	return &Prompt{}
}

func (s *Prompt) GetTemplate(ctx context.Context, userRole int, ask string, history []*schema.Message) ([]*schema.Message, error) {
	template := prompt.FromMessages(schema.FString,
		schema.SystemMessage(common.InitTemplate),
		schema.MessagesPlaceholder("history", false),
	)

	info := map[string]any{
		"role":    common.RoleMap[userRole],
		"ask":     ask,
		"history": history,
	}

	messages, err := template.Format(context.Background(), info)
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Any("ask", ask).Msg("GetTemplate -- 创建提示词模版失败")

		return nil, err
	}

	return messages, nil
}
