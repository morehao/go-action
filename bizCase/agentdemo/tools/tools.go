// Package tools provides InvokableTool implementations used by the ReAct agent.
package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
)

// GetTools returns all available InvokableTool instances for the agent.
func GetTools() []tool.BaseTool {
	return []tool.BaseTool{
		&getCurrentTimeTool{},
		&calculatorTool{},
	}
}

// ─────────────────────────────────────────────
// get_current_time
// ─────────────────────────────────────────────

type getCurrentTimeTool struct{}

func (t *getCurrentTimeTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "get_current_time",
		Desc: "Return the current date and time.",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"format": {
				Type: schema.String,
				Desc: "Optional Go time format string, e.g. \"2006-01-02 15:04:05\". " +
					"Defaults to \"2006-01-02 15:04:05\" when omitted.",
				Required: false,
			},
		}),
	}, nil
}

func (t *getCurrentTimeTool) InvokableRun(_ context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	params := struct {
		Format string `json:"format"`
	}{}
	if err := json.Unmarshal([]byte(argumentsInJSON), &params); err != nil {
		return "", fmt.Errorf("get_current_time: invalid arguments: %w", err)
	}

	format := "2006-01-02 15:04:05"
	if params.Format != "" {
		format = params.Format
	}
	return time.Now().Format(format), nil
}

// ─────────────────────────────────────────────
// calculator
// ─────────────────────────────────────────────

type calculatorTool struct{}

func (t *calculatorTool) Info(_ context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "calculator",
		Desc: "Perform basic arithmetic: add, subtract, multiply, or divide two numbers.",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"operation": {
				Type:     schema.String,
				Desc:     "One of: add, subtract, multiply, divide.",
				Required: true,
			},
			"a": {
				Type:     schema.Number,
				Desc:     "First operand.",
				Required: true,
			},
			"b": {
				Type:     schema.Number,
				Desc:     "Second operand.",
				Required: true,
			},
		}),
	}, nil
}

func (t *calculatorTool) InvokableRun(_ context.Context, argumentsInJSON string, _ ...tool.Option) (string, error) {
	params := struct {
		Operation string  `json:"operation"`
		A         float64 `json:"a"`
		B         float64 `json:"b"`
	}{}
	if err := json.Unmarshal([]byte(argumentsInJSON), &params); err != nil {
		return "", fmt.Errorf("calculator: invalid arguments: %w", err)
	}

	var result float64
	switch params.Operation {
	case "add":
		result = params.A + params.B
	case "subtract":
		result = params.A - params.B
	case "multiply":
		result = params.A * params.B
	case "divide":
		if params.B == 0 {
			return "", fmt.Errorf("calculator: division by zero")
		}
		result = params.A / params.B
	default:
		return "", fmt.Errorf("calculator: unsupported operation %q", params.Operation)
	}

	return strconv.FormatFloat(result, 'f', -1, 64), nil
}
