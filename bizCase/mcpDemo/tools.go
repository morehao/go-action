/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-06-03 12:24:53
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-06-03 12:24:54
 * @FilePath: /go-action/bizCase/mcpDemo/tools.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
