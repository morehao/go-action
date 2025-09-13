package engine

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/model"
)

func routerHuman(ctx context.Context, input string, opts ...any) (output string, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
			state.InterruptFeedback = ""
		}()
		state.Goto = constants.AgentResearchTeam
		if !state.AutoAcceptedPlan {
			switch state.InterruptFeedback {
			case constants.HumanOptionAcceptPlan:
				return nil
			case constants.HumanOptionEditPlan:
				state.Goto = constants.AgentPlanner
				return nil
			default:
				return compose.InterruptAndRerun
			}
		}
		state.Goto = constants.AgentResearchTeam
		return nil
	})
	return output, err
}

func NewHumanNode[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()
	_ = cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(routerHuman))

	_ = cag.AddEdge(compose.START, "router")
	_ = cag.AddEdge("router", compose.END)

	return cag
}
