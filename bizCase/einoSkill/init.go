package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	openaimodel "github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/adk"
	"github.com/cloudwego/eino/adk/middlewares/skill"
)

var agentRunner *adk.Runner

func init() {
	ctx := context.Background()

	// Determine skills directory: env var takes precedence, then fallback to
	// the source-relative path (convenient for `go run .` during development).
	skillsDir := os.Getenv("SKILLS_DIR")
	if skillsDir == "" {
		_, srcFile, _, ok := runtime.Caller(0)
		if !ok {
			panic("failed to get caller info")
		}
		skillsDir = filepath.Join(filepath.Dir(srcFile), "skills")
	}

	cm, err := openaimodel.NewChatModel(ctx, &openaimodel.ChatModelConfig{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  getEnvOrDefault("OPENAI_MODEL", "gpt-4o-mini"),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create chat model: %v", err))
	}

	backend := newLocalSkillBackend(skillsDir)

	sm, err := skill.NewMiddleware(ctx, &skill.Config{
		Backend: backend,
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create skill middleware: %v", err))
	}

	agent, err := adk.NewChatModelAgent(ctx, &adk.ChatModelAgentConfig{
		Name:        "SkillAgent",
		Description: "A helpful programming assistant with domain knowledge skills.",
		Instruction: "You are a helpful programming assistant. " +
			"Use the available skills to provide detailed and accurate guidance.",
		Model:    cm,
		Handlers: []adk.ChatModelAgentMiddleware{sm},
	})
	if err != nil {
		panic(fmt.Sprintf("failed to create agent: %v", err))
	}

	agentRunner = adk.NewRunner(ctx, adk.RunnerConfig{
		Agent:           agent,
		EnableStreaming: true,
	})
}

func getEnvOrDefault(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}
