// Package agent provides a ReAct agent built with the Eino framework.
package agent

import (
	"context"
	"fmt"
	"sync"

	"github.com/cloudwego/eino-ext/components/model/deepseek"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/morehao/go-action/bizcase/agentdemo/config"
	"github.com/morehao/go-action/bizcase/agentdemo/tools"
)

var (
	once       sync.Once
	reactAgent *react.Agent
)

// Init creates and stores the global ReAct agent singleton.
// It must be called once before any handler processes requests.
func Init(ctx context.Context) error {
	var initErr error
	once.Do(func() {
		chatModel, err := deepseek.NewChatModel(ctx, &deepseek.ChatModelConfig{
			APIKey:  config.Cfg.Model.APIKey,
			Model:   config.Cfg.Model.Model,
			BaseURL: config.Cfg.Model.BaseURL,
		})
		if err != nil {
			initErr = fmt.Errorf("agentdemo: create chat model: %w", err)
			return
		}

		reactAgent, err = react.NewAgent(ctx, &react.AgentConfig{
			ToolCallingModel: chatModel,
			ToolsConfig: compose.ToolsNodeConfig{
				Tools: tools.GetTools(),
			},
			MessageModifier: react.NewPersonaModifier(
				"You are a helpful assistant. " +
					"Use the provided tools whenever they can help answer the user's question.",
			),
		})
		if err != nil {
			initErr = fmt.Errorf("agentdemo: create react agent: %w", err)
		}
	})
	return initErr
}

// GetAgent returns the initialized ReAct agent.
func GetAgent() *react.Agent {
	return reactAgent
}
