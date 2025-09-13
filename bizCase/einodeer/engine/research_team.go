package engine

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func routerResearchTeam(ctx context.Context, input string, opts ...any) (output string, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
		}()
		state.Goto = constants.AgentPlanner
		if state.CurrentPlan == nil {
			return nil
		}
		for i, step := range state.CurrentPlan.Steps {
			if step.ExecutionRes != nil {
				continue
			}
			glog.Infof(ctx, "research_team_step: %v, index: %d", step, i)
			switch step.StepType {
			case model.Research:
				state.Goto = constants.AgentResearcher
				return nil
			case model.Processing:
				state.Goto = constants.AgentCoder
				return nil
			}
		}
		if state.PlanIterations >= state.MaxPlanIterations {
			state.Goto = constants.AgentReporter
			return nil
		}
		return nil
	})
	return output, nil
}

func NewResearchTeamNode[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()
	_ = cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(routerResearchTeam))

	_ = cag.AddEdge(compose.START, "router")
	_ = cag.AddEdge("router", compose.END)

	return cag
}
