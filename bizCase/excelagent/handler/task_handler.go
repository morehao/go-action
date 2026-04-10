package handler

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"excelagent/config"
	"excelagent/service/agent"
	"excelagent/types"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TaskHandler struct {
	cfg *config.Config
}

func NewTaskHandler(cfg *config.Config) *TaskHandler {
	return &TaskHandler{
		cfg: cfg,
	}
}

func (h *TaskHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/tasks", h.CreateTask)
	}
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var req types.TaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, types.ErrorResponse{
			Error:   "invalid_request",
			Message: err.Error(),
		})
		return
	}

	workDir, err := agent.PrepareWorkDir(h.cfg, req.Files)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:   "workdir_error",
			Message: err.Error(),
		})
		return
	}

	task := &types.Task{
		ID:        uuid.New().String(),
		Query:     req.Query,
		Files:     req.Files,
		WorkDir:   workDir,
		CreatedAt: time.Now(),
	}

	ctx := context.Background()
	a, err := agent.NewAgent(ctx, h.cfg, task.WorkDir)
	if err != nil {
		c.JSON(http.StatusInternalServerError, types.ErrorResponse{
			Error:   "agent_error",
			Message: err.Error(),
		})
		return
	}

	if err := a.Run(ctx, task); err != nil {
		report := h.generateReport(task)
		c.JSON(http.StatusOK, types.TaskResultResponse{
			TaskID:      task.ID,
			Status:      types.TaskStatusFailed,
			Result:      err.Error(),
			OutputFiles: h.collectOutputFiles(task.WorkDir),
			Report:      report,
			CompletedAt: time.Now(),
		})
		return
	}

	outputFiles := h.collectOutputFiles(task.WorkDir)
	report := h.generateReportWithFiles(task, outputFiles)

	c.JSON(http.StatusOK, types.TaskResultResponse{
		TaskID:      task.ID,
		Status:      types.TaskStatusCompleted,
		Result:      task.Result,
		OutputFiles: outputFiles,
		Report:      report,
		CompletedAt: time.Now(),
	})
}

func (h *TaskHandler) collectOutputFiles(workDir string) []string {
	var files []string
	entries, _ := filepath.Glob(filepath.Join(workDir, "*"))
	for _, e := range entries {
		if !strings.HasSuffix(e, ".py") {
			files = append(files, filepath.Base(e))
		}
	}
	return files
}

func (h *TaskHandler) generateReport(task *types.Task) string {
	var sb strings.Builder
	sb.WriteString("## Task Report\n\n")
	sb.WriteString("### Result\n")
	sb.WriteString(task.Result)
	return sb.String()
}

func (h *TaskHandler) generateReportWithFiles(task *types.Task, outputFiles []string) string {
	var sb strings.Builder
	sb.WriteString("## Task Report\n\n")
	sb.WriteString("### Result\n")
	sb.WriteString(task.Result)
	sb.WriteString("\n\n### Output Files\n")
	for _, f := range outputFiles {
		sb.WriteString("- " + f + "\n")
	}
	return sb.String()
}
