package agentchat

import "context"

type Planner interface {
	PlanTasks(ctx context.Context) ([]Task, error)
	GetPlannerType() string
	GetPlanHistory(ctx context.Context, QuestionID string) ([]Task, error)
}

type Storage interface {
}

type planContext interface {
	Storage
	GetPlanInput(ctx context.Context, QuestionID string) (string, error)
	GetPlanOutput(ctx context.Context, QuestionID string) (string, error)
	GetInputByAgentName(ctx context.Context, agentName string) (string, error)
	GetOutputByAgentName(ctx context.Context, agentName string) (string, error)
}

type SettingPlan struct{}

type LLMPlan struct{}

type GraphPlan struct{}
