package services

import (
	"ai-ZhenLou/global"
	"context"
	"github.com/cloudwego/eino-ext/components/model/ollama"
)

type ChatModel struct {
}

func NewChatModel() *ChatModel {
	return &ChatModel{}
}

func (s *ChatModel) LocalOllama(ctx context.Context) (*ollama.ChatModel, error) {
	chatModel, err := ollama.NewChatModel(ctx, &ollama.ChatModelConfig{
		BaseURL: "http://localhost:11434", // Ollama 服务地址
		Model:   "qwen3:1.7b",             // 模型名称
	})

	if err != nil {
		global.App.Log.Err(err).Msg("LocalModel -- 创建模型失败")
		return nil, err
	}

	return chatModel, nil
}
