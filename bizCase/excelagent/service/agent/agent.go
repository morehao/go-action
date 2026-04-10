package agent

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"excelagent/config"
	"excelagent/types"

	"github.com/cloudwego/eino-ext/components/model/ark"
	"github.com/cloudwego/eino-ext/components/tool/commandline"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/prebuilt/planexecute"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

type Agent struct {
	cfg      *config.Config
	operator *LocalOperator
	agent    adk.Agent
}

type LocalOperator struct {
	workDir string
}

func (l *LocalOperator) ReadFile(ctx context.Context, path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return err.Error(), nil
	}
	return string(b), nil
}

func (l *LocalOperator) WriteFile(ctx context.Context, path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

func (l *LocalOperator) IsDirectory(ctx context.Context, path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func (l *LocalOperator) Exists(ctx context.Context, path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (l *LocalOperator) RunCommand(ctx context.Context, command []string) (*commandline.CommandOutput, error) {
	cmd := exec.CommandContext(ctx, command[0], command[1:]...)
	cmd.Dir = l.workDir

	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("command failed: %v, output: %s", err, string(out))
	}
	return &commandline.CommandOutput{
		Stdout: string(out),
	}, nil
}

func NewAgent(ctx context.Context, cfg *config.Config, workDir string) (*Agent, error) {
	operator := &LocalOperator{workDir: workDir}

	chatModel, err := newChatModel(ctx, cfg)
	if err != nil {
		return nil, err
	}

	p, err := newPlanner(ctx, chatModel)
	if err != nil {
		return nil, err
	}

	e, err := newExecutor(ctx, operator)
	if err != nil {
		return nil, err
	}

	rp, err := newReplanner(ctx, chatModel)
	if err != nil {
		return nil, err
	}

	planExecuteAgent, err := planexecute.New(ctx, &planexecute.Config{
		Planner:       p,
		Executor:      e,
		Replanner:     rp,
		MaxIterations: cfg.Task.MaxIterations,
	})
	if err != nil {
		return nil, err
	}

	return &Agent{
		cfg:      cfg,
		operator: operator,
		agent:    planExecuteAgent,
	}, nil
}

func newChatModel(ctx context.Context, cfg *config.Config) (model.ToolCallingChatModel, error) {
	cm, err := ark.NewChatModel(ctx, &ark.ChatModelConfig{
		APIKey:  cfg.Eino.APIKey,
		BaseURL: cfg.Eino.BaseURL,
		Model:   cfg.Eino.Model,
	})
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func newPlanner(ctx context.Context, cm model.ToolCallingChatModel) (adk.Agent, error) {
	return planexecute.NewPlanner(ctx, &planexecute.PlannerConfig{
		ChatModelWithFormattedOutput: cm,
	})
}

func newExecutor(ctx context.Context, op *LocalOperator) (adk.Agent, error) {
	return planexecute.NewExecutor(ctx, &planexecute.ExecutorConfig{
		MaxIterations: 20,
	})
}

func newReplanner(ctx context.Context, cm model.ToolCallingChatModel) (adk.Agent, error) {
	return planexecute.NewReplanner(ctx, &planexecute.ReplannerConfig{
		ChatModel: cm,
	})
}

func (a *Agent) Run(ctx context.Context, task *types.Task) error {
	runner := adk.NewRunner(ctx, adk.RunnerConfig{
		Agent:           a.agent,
		EnableStreaming: false,
	})

	query := schema.UserMessage(task.Query)
	iter := runner.Run(ctx, []*schema.Message{query})

	for {
		event, ok := iter.Next()
		if !ok {
			break
		}
		if event.Output != nil && event.Output.MessageOutput != nil {
			msg := event.Output.MessageOutput.Message
			if msg != nil {
				task.Result = msg.String()
			}
		}
	}

	return nil
}

func PrepareWorkDir(cfg *config.Config, files []types.File) (string, error) {
	workDir := filepath.Join(cfg.Task.Workspace, fmt.Sprintf("task_%d", os.Getpid()))
	if err := os.MkdirAll(workDir, 0755); err != nil {
		return "", err
	}

	for _, f := range files {
		srcPath := filepath.Join(cfg.Task.Workspace, "input", f.Name)
		dstPath := filepath.Join(workDir, f.Name)
		data, err := os.ReadFile(srcPath)
		if err != nil {
			continue
		}
		if err := os.WriteFile(dstPath, data, 0644); err != nil {
			continue
		}
	}

	return workDir, nil
}
