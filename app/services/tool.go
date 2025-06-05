package services

import (
	"ai-ZhenLou/global"
	"context"
	mcpp "github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/components/tool"
	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
}

func NewTool() *Tool {
	return &Tool{}
}

func (s *Tool) McpOfMysql(ctx context.Context) ([]tool.BaseTool, error) {
	cli, err := client.NewSSEMCPClient("http://localhost:12345/sse")
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Msg("McpOfMysql -- mysql的mcp服务创建cli失败")
		return nil, err
	}

	err = cli.Start(ctx)
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Msg("McpOfMysql -- mysql的mcp服务创建cli启动失败")
		return nil, err
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "mysql-mcp-client",
		Version: "1.0.0",
	}

	_, err = cli.Initialize(ctx, initRequest)
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Msg("McpOfMysql -- mysql的mcp服务初始化失败")
		return nil, err
	}

	tools, err := mcpp.GetTools(ctx, &mcpp.Config{Cli: cli})
	if err != nil {
		// 打印日志
		global.App.Log.Err(err).Msg("McpOfMysql -- mysql的mcp服务初始化失败")
		return nil, err
	}

	return tools, nil
}
