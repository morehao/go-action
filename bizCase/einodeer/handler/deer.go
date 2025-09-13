package handler

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/compose"
	"github.com/gin-gonic/gin"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/engine"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/go-action/bizCase/einodeer/utils"
	"github.com/morehao/golib/gcontext/gincontext"
	"github.com/morehao/golib/glog"
)

func ChatStreamEino(ctx *gin.Context) {
	// 设置响应头（NewStream 会自动设置部分头，但建议显式声明）
	ctx.Header("Content-Type", "text/event-stream; charset=utf-8")
	ctx.Header("Cache-Control", "no-cache")
	ctx.Header("Connection", "keep-alive")
	ctx.Header("Access-Control-Allow-Origin", "*")

	// 请求体校验
	req := new(model.ChatRequest)
	if err := ctx.ShouldBind(&req); err != nil {
		glog.Errorf(ctx, "ChatStreamEino_bind_error: %v", err)
		gincontext.Fail(ctx, err)
		return
	}
	glog.Infof(ctx, "ChatStream_begin: %s", glog.ToJsonString(req))

	// 根据前端参数生成Graph State
	genFunc := func(ctx context.Context) *model.State {
		return &model.State{
			MaxPlanIterations:             req.MaxPlanIterations,
			MaxStepNum:                    req.MaxStepNum,
			Messages:                      req.Messages,
			Goto:                          constants.AgentCoordinator,
			EnableBackgroundInvestigation: req.EnableBackgroundInvestigation,
		}
	}

	// Build Graph
	r := engine.Builder[string, string, *model.State](ctx, genFunc)

	// Run Graph
	_, err := r.Stream(ctx, constants.AgentCoordinator,
		compose.WithCheckPointID(req.ThreadID), // 指定Graph的CheckPointID
		// 中断后，获取用户的edit_plan信息
		compose.WithStateModifier(func(ctx context.Context, path compose.NodePath, state any) error {
			s := state.(*model.State)
			s.InterruptFeedback = req.InterruptFeedback
			if req.InterruptFeedback == "edit_plan" {
				s.Messages = append(s.Messages, req.Messages...)
			}
			glog.Debugf(ctx, "ChatStream_modf, path: %s, state: %s", path.GetPath(), glog.ToJsonString(state))
			return nil
		}),
		// 连接LoggerCallback
		compose.WithCallbacks(&infra.LoggerCallback{
			ID:  req.ThreadID,
			SSE: ctx.Writer,
		}),
	)

	// 将interrupt信号传递到前端
	if info, ok := compose.ExtractInterruptInfo(err); ok {
		glog.Debugf(ctx, "ChatStream_interrupt: %s", glog.ToJsonString(info))
		data := &model.ChatResp{
			ThreadID:     req.ThreadID,
			ID:           "human_feedback:" + utils.RandStr(20),
			Role:         "assistant",
			Content:      "检查计划",
			FinishReason: "interrupt",
			Options: []map[string]any{
				{
					"text":  "编辑计划",
					"value": "edit_plan",
				},
				{
					"text":  "开始执行",
					"value": "accepted",
				},
			},
		}
		dB, _ := json.Marshal(data)
		if err := infra.WriteSSE(ctx.Writer, "", "interrupt", dB); err != nil {
			glog.Errorf(ctx, "ChatStream_interrupt: %v", err)
		}
	}
	if err != nil {
		glog.Errorf(ctx, "ChatStream_error: %v", err)
	}
}
