package infra

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"github.com/morehao/go-action/bizCase/einodeer/config"
	"github.com/morehao/golib/glog"
)

// ToolManager 管理Eino工具
type ToolManager struct {
	tools map[string]tool.InvokableTool
}

// NewToolManager 创建工具管理器
func NewToolManager() *ToolManager {
	return &ToolManager{
		tools: make(map[string]tool.InvokableTool),
	}
}

// RegisterTool 注册工具
func (tm *ToolManager) RegisterTool(name string, t tool.InvokableTool) {
	tm.tools[name] = t
}

// GetTool 获取工具
func (tm *ToolManager) GetTool(name string) tool.InvokableTool {
	return tm.tools[name]
}

// GetToolByNameSuffix 通过名称后缀获取工具
func (tm *ToolManager) GetToolByNameSuffix(suffix string) tool.InvokableTool {
	for name, t := range tm.tools {
		if strings.HasSuffix(name, suffix) {
			return t
		}
	}
	return nil
}

// GetAllTools 获取所有工具
func (tm *ToolManager) GetAllTools() map[string]tool.InvokableTool {
	return tm.tools
}

// 全局工具管理器实例
var DefaultToolManager = NewToolManager()

// InitTools 初始化所有工具
func InitTools() error {
	// 初始化搜索工具
	for name, server := range config.Config.Tools.Servers {
		if err := initToolFromConfig(name, server); err != nil {
			return err
		}
	}

	return nil
}

// initToolFromConfig 从配置初始化工具
func initToolFromConfig(name string, server struct {
	Command string            `yaml:"command"`
	Args    []string          `yaml:"args"`
	Env     map[string]string `yaml:"env,omitempty"`
}) error {
	// 这里简化实现，根据配置创建相应的工具
	// 在实际应用中，可能需要更复杂的逻辑来创建不同类型的工具
	glog.Infof(context.Background(), "initializing tool: %s", name)

	// 创建一个简单的工具实现
	tool := &ExternalProcessTool{
		name:    name,
		command: server.Command,
		args:    server.Args,
		env:     server.Env,
	}

	// 注册工具
	DefaultToolManager.RegisterTool(name, tool)
	return nil
}

// ExternalProcessTool 外部进程工具实现
type ExternalProcessTool struct {
	name    string
	command string
	args    []string
	env     map[string]string
}

// Info 返回工具信息
func (t *ExternalProcessTool) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: t.name,
		Desc: fmt.Sprintf("External process tool: %s", t.command),
	}, nil
}

// InvokableRun 执行工具
func (t *ExternalProcessTool) InvokableRun(ctx context.Context, input string, opts ...tool.Option) (string, error) {
	// 创建命令
	cmd := exec.CommandContext(ctx, t.command, append(t.args, input)...)

	// 设置环境变量
	if t.env != nil {
		env := cmd.Environ()
		for k, v := range t.env {
			env = append(env, fmt.Sprintf("%s=%s", k, v))
		}
		cmd.Env = env
	}

	// 执行命令并获取输出
	output, err := cmd.CombinedOutput()
	if err != nil {
		glog.Errorf(ctx, "tool execution error: %v, output: %s", err, string(output))
		return "", err
	}

	return string(output), nil
}
