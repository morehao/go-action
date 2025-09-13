package engine

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cloudwego/eino-examples/flow/agent/deer-go/biz/infra"
	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func loadMsg(ctx context.Context, name string, opts ...any) (output []*schema.Message, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		sysPrompt, err := infra.GetPromptTemplate(ctx, name)
		if err != nil {
			glog.Errorf(ctx, "get prompt template error: %v", err)
			return err
		}

		promptTemp := prompt.FromMessages(schema.Jinja2,
			schema.SystemMessage(sysPrompt),
			schema.MessagesPlaceholder("user_input", true),
		)

		variables := map[string]any{
			"locale":              state.Locale,
			"max_step_num":        state.MaxStepNum,
			"max_plan_iterations": state.MaxPlanIterations,
			"CURRENT_TIME":        time.Now().Format("2006-01-02 15:04:05"),
			"user_input":          state.Messages,
		}
		output, err = promptTemp.Format(ctx, variables)
		return err
	})
	return output, err
}

func router(ctx context.Context, input *schema.Message, opts ...any) (output string, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
		}()
		state.Goto = compose.END
		if len(input.ToolCalls) > 0 && input.ToolCalls[0].Function.Name == "hand_to_planner" {
			argMap := map[string]string{}
			_ = json.Unmarshal([]byte(input.ToolCalls[0].Function.Arguments), &argMap)
			state.Locale, _ = argMap["locale"]
			if state.EnableBackgroundInvestigation {
				state.Goto = constants.AgentBackgroundInvestigator
			} else {
				state.Goto = constants.AgentPlanner
			}
		}
		return nil
	})
	return output, nil
}

func NewCAgent[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()

	hand_to_planner := &schema.ToolInfo{
		Name: "hand_to_planner",
		Desc: "Handoff to planner agent to do plan.",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"task_title": {
				Type:     schema.String,
				Desc:     "The title of the task to be handed off.",
				Required: true,
			},
			"locale": {
				Type:     schema.String,
				Desc:     "The user's detected language locale (e.g., en-US, zh-CN).",
				Required: true,
			},
		}),
	}

	coorModel, _ := infra.ChatModel.WithTools([]*schema.ToolInfo{hand_to_planner})

	_ = cag.AddLambdaNode("load", compose.InvokableLambdaWithOption(loadMsg))
	_ = cag.AddChatModelNode("agent", coorModel)
	_ = cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(router))

	_ = cag.AddEdge(compose.START, "load")
	_ = cag.AddEdge("load", "agent")
	_ = cag.AddEdge("agent", "router")
	_ = cag.AddEdge("router", compose.END)
	return cag
}
