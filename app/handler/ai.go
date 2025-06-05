package handler

import (
	"ai-ZhenLou/app/req"
	"ai-ZhenLou/app/resp"
	"ai-ZhenLou/app/services"
	"ai-ZhenLou/global"
	"ai-ZhenLou/utils"
	"fmt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/gin-gonic/gin"
)

type AI struct {
	MessageSvc   *services.Message
	UserSvc      *services.User
	PromptSvc    *services.Prompt
	ChatModelSvc *services.ChatModel
	ToolSvc      *services.Tool
}

func NewAI(msgSvc *services.Message, userSvc *services.User, promptSvc *services.Prompt, chatModelSvc *services.ChatModel, toolSvc *services.Tool) *AI {
	return &AI{
		MessageSvc:   msgSvc,
		UserSvc:      userSvc,
		PromptSvc:    promptSvc,
		ChatModelSvc: chatModelSvc,
		ToolSvc:      toolSvc,
	}
}

func (h *AI) Talk(c *gin.Context) {
	var (
		r   req.TalkReq
		ctx = c.Request.Context()
	)

	if err := c.ShouldBindJSON(&r); err != nil {
		resp.Fail(c, global.ParamErrCode, global.ParamErrMsg)
		return
	}

	if err := r.Check(); err != nil {
		resp.Fail(c, global.ParamErrCode, err.Error())
		return
	}

	// step1. 获取用户信息与历史聊天
	userInfo, err := h.UserSvc.GetUser(ctx, r.UserID)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}
	// 如果没有会话id，则表示第一次会话，不用查询，但要生成一个新的会话id
	// 有会话id，则调用message服务获取
	historyMessage, err := h.MessageSvc.GetMessageHistory(ctx, r.UserID, r.DialogueID)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}

	if len(r.DialogueID) == 0 {
		r.DialogueID = utils.CreateKeyByTime(r.UserID)
	}

	// step2. 构建提示词模版
	prompt, err := h.PromptSvc.GetTemplate(ctx, userInfo.Role, r.Ask, historyMessage)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}
	// step3. 构建本地大模型
	chatModel, err := h.ChatModelSvc.LocalOllama(ctx)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}
	// step4. 构建工具集
	tools, err := h.ToolSvc.McpOfMysql(ctx)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, err.Error())
		return
	}
	// step5. 构造agent
	agent, err := react.NewAgent(ctx, &react.AgentConfig{
		ToolCallingModel: chatModel,
		ToolsConfig: compose.ToolsNodeConfig{
			Tools: tools,
		},
	})
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, fmt.Errorf("构造agent错误；%s", err.Error()).Error())
		return
	}
	// step6. 执行agent并获取结果
	outMessage, err := agent.Generate(ctx, prompt)
	if err != nil {
		resp.Fail(c, global.ProcessErrCode, fmt.Errorf("执行agent错误；%s", err.Error()).Error())
		return
	}
	// step7. 更新用户的聊天历史
	if err = h.MessageSvc.UpdMessageHistory(ctx, r.Ask, outMessage, r.UserID, r.DialogueID); err != nil {
		resp.Fail(c, global.ProcessErrCode, fmt.Errorf("保存本次记录错误；%s", err.Error()).Error())
		return
	}

	resp.Success(c, outMessage.String())
}
