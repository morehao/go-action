package engine

import (
	"context"

	"github.com/cloudwego/eino/compose"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

//type I = string
//type O = string

// 子图流转函数，由上一个子图决定接下来流转到哪个agent
// 并将其name写入 state.Goto ，该函数读取 state.Goto 并将控制权交给对应agent
func agentHandOff(ctx context.Context, input string) (next string, err error) {
	defer func() {
		glog.Infof(ctx, "agent_hand_off, input: %s, next: %s", input, next)
	}()
	_ = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		next = state.Goto
		return nil
	})
	return next, nil
}

// Builder 初始化全部子图并连接
func Builder[I, O, S any](ctx context.Context, genFunc compose.GenLocalState[S]) compose.Runnable[I, O] {

	g := compose.NewGraph[I, O](
		compose.WithGenLocalState(genFunc),
	)

	outMap := map[string]bool{
		constants.AgentCoordinator:            true,
		constants.AgentPlanner:                true,
		constants.AgentReporter:               true,
		constants.AgentResearchTeam:           true,
		constants.AgentResearcher:             true,
		constants.AgentCoder:                  true,
		constants.AgentBackgroundInvestigator: true,
		constants.AgentHuman:                  true,
		compose.END:                           true,
	}

	coordinatorGraph := NewCAgent[I, O](ctx)
	plannerGraph := NewPlanner[I, O](ctx)
	reporterGraph := NewReporter[I, O](ctx)
	researchTeamGraph := NewResearchTeamNode[I, O](ctx)
	researcherGraph := NewResearcher[I, O](ctx)
	bIGraph := NewBAgent[I, O](ctx)
	coder := NewCoder[I, O](ctx)
	human := NewHumanNode[I, O](ctx)

	_ = g.AddGraphNode(constants.AgentCoordinator, coordinatorGraph, compose.WithNodeName(constants.AgentCoordinator))
	_ = g.AddGraphNode(constants.AgentPlanner, plannerGraph, compose.WithNodeName(constants.AgentPlanner))
	_ = g.AddGraphNode(constants.AgentReporter, reporterGraph, compose.WithNodeName(constants.AgentReporter))
	_ = g.AddGraphNode(constants.AgentResearchTeam, researchTeamGraph, compose.WithNodeName(constants.AgentResearchTeam))
	_ = g.AddGraphNode(constants.AgentResearcher, researcherGraph, compose.WithNodeName(constants.AgentResearcher))
	_ = g.AddGraphNode(constants.AgentCoder, coder, compose.WithNodeName(constants.AgentCoder))
	_ = g.AddGraphNode(constants.AgentBackgroundInvestigator, bIGraph, compose.WithNodeName(constants.AgentBackgroundInvestigator))
	_ = g.AddGraphNode(constants.AgentHuman, human, compose.WithNodeName(constants.AgentHuman))

	_ = g.AddBranch(constants.AgentCoordinator, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentPlanner, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentReporter, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentResearchTeam, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentResearcher, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentCoder, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentBackgroundInvestigator, compose.NewGraphBranch(agentHandOff, outMap))
	_ = g.AddBranch(constants.AgentHuman, compose.NewGraphBranch(agentHandOff, outMap))

	_ = g.AddEdge(compose.START, constants.AgentCoordinator)

	r, err := g.Compile(ctx,
		compose.WithGraphName("EinoDeer"),
		compose.WithNodeTriggerMode(compose.AnyPredecessor),
		compose.WithCheckPointStore(model.NewDeerCheckPoint(ctx)), // 指定Graph CheckPointStore
	)
	if err != nil {
		glog.Errorf(ctx, "compile failed: %v", err)
	}
	return r
}
