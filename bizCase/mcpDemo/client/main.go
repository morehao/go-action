package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/morehao/golib/gutils"
)

func main() {
	// create a new MCP client
	client, newClientErr := client.NewStdioMCPClient(
		"../server/mcp-server",
		[]string{},
	)
	if newClientErr != nil {
		panic(newClientErr)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	initRequest := mcp.InitializeRequest{
		Params: mcp.InitializeParams{
			ProtocolVersion: mcp.LATEST_PROTOCOL_VERSION,
			ClientInfo: mcp.Implementation{
				Name:    "Client Demo",
				Version: "1.0.0",
			},
		},
	}

	initResult, initResultErr := client.Initialize(ctx, initRequest)
	if initResultErr != nil {
		panic(initResultErr)
	}

	fmt.Println("Initialization successful, server info :", initResult.ServerInfo.Name, initResult.ServerInfo.Version)

	fmt.Println("prompt list")
	promptsRequest := mcp.ListPromptsRequest{}
	promptsResult, promptsResultErr := client.ListPrompts(ctx, promptsRequest)
	if promptsResultErr != nil {
		panic(promptsResultErr)
	}
	for _, v := range promptsResult.Prompts {
		fmt.Printf("Prompt info: %s, Description: %s, Arguments: %s\n", v.Name, v.Description, gutils.ToJsonString(v.Arguments))
	}

	fmt.Println("resource list")
	resourcesRequest := mcp.ListResourcesRequest{}
	resourcesResult, resourcesResultErr := client.ListResources(ctx, resourcesRequest)
	if resourcesResultErr != nil {
		panic(resourcesResultErr)
	}
	for _, v := range resourcesResult.Resources {
		fmt.Printf("Resource info, uri: %s, name: %s, description: %s, MIME类型: %s\n", v.URI, v.Name, v.Description, v.MIMEType)
	}

	fmt.Println("tool list")
	toolsRequest := mcp.ListToolsRequest{}
	toolsResult, toolsResultErr := client.ListTools(ctx, toolsRequest)
	if toolsResultErr != nil {
		panic(toolsResultErr)
	}
	for _, v := range toolsResult.Tools {
		fmt.Printf("Tool info, name: %s, Description: %s, Arguments: %s\n", v.Name, v.Description, gutils.ToJsonString(v.InputSchema.Properties))
	}

	fmt.Println("tool call")
	toolRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	toolRequest.Params.Name = "calculate"
	toolRequest.Params.Arguments = map[string]any{
		"operation": "add",
		"x":         1,
		"y":         1,
	}
	callToolResult, callToolResultErr := client.CallTool(ctx, toolRequest)
	if callToolResultErr != nil {
		panic(callToolResultErr)
	}
	fmt.Println("Tool call result:", callToolResult.Content[0].(mcp.TextContent).Text)

}
