package types

import "time"

type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusProcessing TaskStatus = "processing"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusFailed     TaskStatus = "failed"
)

type File struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type TaskOptions struct {
	MaxIterations int `json:"max_iterations,omitempty"`
	Timeout       int `json:"timeout,omitempty"`
}

type TaskRequest struct {
	Query   string       `json:"query" binding:"required"`
	Files   []File       `json:"files"`
	Options *TaskOptions `json:"options,omitempty"`
}

type TaskResponse struct {
	TaskID    string     `json:"task_id"`
	Status    TaskStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
}

type TaskResultResponse struct {
	TaskID      string     `json:"task_id"`
	Status      TaskStatus `json:"status"`
	Result      string     `json:"result"`
	OutputFiles []string   `json:"output_files"`
	Report      string     `json:"report"`
	CompletedAt time.Time  `json:"completed_at"`
}

type Task struct {
	ID          string
	Query       string
	Files       []File
	Result      string
	OutputFiles []string
	CreatedAt   time.Time
	WorkDir     string
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}
