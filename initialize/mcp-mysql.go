package initialize

import (
	"ai-ZhenLou/global"
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func MysqlMcpStart() {
	svr := server.NewMCPServer("mysql-mcp-server", mcp.LATEST_PROTOCOL_VERSION)
	svr.AddTool(mcp.NewTool("mysql-mcp",
		mcp.WithDescription("mysql的工具，可以查询mysql的数据，包括各种表"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("要执行的sql语句"),
		),
	), func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		arg := request.Params.Arguments.(map[string]any)
		query := arg["query"].(string)
		// 执行语句
		rows, err := global.App.DB.Raw(query).Rows()
		if err != nil {
			// 打印日志
			global.App.Log.Err(err).Any("sql", query).Msg("mysql-mcp -- 执行语句失败")
			return nil, err
		}
		defer rows.Close()
		// 获取列名
		columns, err := rows.Columns()
		if err != nil {
			// 打印日志
			global.App.Log.Err(err).Any("sql", query).Msg("mysql-mcp -- 获取列名失败")
			return nil, err
		}
		// 遍历结果
		var result string
		for rows.Next() {
			values := make([]interface{}, len(columns))
			for i := range values {
				values[i] = new(interface{})
			}
			if err = rows.Scan(values...); err != nil {
				global.App.Log.Err(err).Any("sql", query).Msg("mysql-mcp -- 获取结果失败")
				return nil, err
			}
			for i, val := range values {
				// 获取实际值
				tmpI := *val.(*interface{})
				// 断言类型
				switch tmpI.(type) {
				case string:
					result += fmt.Sprintf("%s: %v\n", columns[i], tmpI.(string))
				case []byte:
					result += fmt.Sprintf("%s: %v\n", columns[i], string(tmpI.([]byte)))
				default:
					result += fmt.Sprintf("%s: %v\n", columns[i], tmpI)
				}
			}
			result += "----\n"
		}

		if result == "" {
			result = "查询成功，但没有返回数据"
		}

		return mcp.NewToolResultText(result), nil
	})
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				fmt.Println(e)
			}
		}()

		err := server.NewSSEServer(svr, server.WithBaseURL("http://localhost:12345")).Start("localhost:12345")

		if err != nil {
			fmt.Println("mysql-mcp-server 启动失败")
			global.App.Log.Err(err).Msg("mysql-mcp-server 启动失败")
			return
		}

		fmt.Println("mysql-mcp-server 启动成功")
		global.App.Log.Err(err).Msg("mysql-mcp-server 启动成功")
	}()
}
