package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/schema"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func loadPlannerMsg(ctx context.Context, name string, opts ...any) (output []*schema.Message, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		sysPrompt, err := infra.GetPromptTemplate(ctx, name)
		if err != nil {
			glog.Errorf(ctx, "get prompt template error: %v", err)
			return err
		}

		var promptTemp *prompt.DefaultChatTemplate
		if state.EnableBackgroundInvestigation && len(state.BackgroundInvestigationResults) > 0 {
			promptTemp = prompt.FromMessages(schema.Jinja2,
				schema.SystemMessage(sysPrompt),
				schema.MessagesPlaceholder("user_input", true),
				schema.UserMessage(fmt.Sprintf("background investigation results of user query: \n %s", state.BackgroundInvestigationResults)),
			)
		} else {
			promptTemp = prompt.FromMessages(schema.Jinja2,
				schema.SystemMessage(sysPrompt),
				schema.MessagesPlaceholder("user_input", true),
			)
		}

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

func routerPlanner(ctx context.Context, input *schema.Message, opts ...any) (output string, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
		}()
		state.Goto = compose.END
		state.CurrentPlan = &model.Plan{}
		// TODO fix 一些 ```
		err = json.Unmarshal([]byte(input.Content), state.CurrentPlan)
		if err != nil {
			glog.Errorf(ctx, "gen_plan_fail, error: %v, input.Content: %s", err, input.Content)
			if state.PlanIterations > 0 {
				state.Goto = constants.AgentReporter
				return nil
			}
			return nil
		}
		glog.Infof(ctx, "gen_plan_ok, plan: %v", state.CurrentPlan)
		state.PlanIterations++
		if state.CurrentPlan.HasEnoughContext {
			state.Goto = constants.AgentReporter
			return nil
		}

		state.Goto = constants.AgentHuman // TODO 改成 human_feedback
		return nil
	})
	return output, nil
}

func NewPlanner[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()

	_ = cag.AddLambdaNode("load", compose.InvokableLambdaWithOption(loadPlannerMsg))
	_ = cag.AddChatModelNode("agent", infra.PlanModel)
	_ = cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(routerPlanner))

	_ = cag.AddEdge(compose.START, "load")
	_ = cag.AddEdge("load", "agent")
	_ = cag.AddEdge("agent", "router")
	_ = cag.AddEdge("router", compose.END)
	return cag
}
