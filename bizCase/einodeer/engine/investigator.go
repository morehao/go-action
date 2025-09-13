package engine

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/cloudwego/eino-ext/components/tool/mcp"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func search(ctx context.Context, name string, opts ...any) (output string, err error) {
	var searchTool tool.InvokableTool
	for _, cli := range infra.MCPServer {
		if searchTool != nil {
			break
		}
		ts, err := mcp.GetTools(ctx, &mcp.Config{Cli: cli})
		if err != nil {
			glog.Errorf(ctx, "get tools error: %v", err)
			continue
		}
		for _, t := range ts {
			info, _ := t.Info(ctx)
			if strings.HasSuffix(info.Name, "search") {
				searchTool, _ = t.(tool.InvokableTool)
				break
			}
		}
	}

	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		args := map[string]any{
			"query": state.Messages[len(state.Messages)-1].Content,
		}
		argsBytes, err := json.Marshal(args)
		if err != nil {
			glog.Errorf(ctx, "json marshal error: %v", err)
			return err
		}
		result, err := searchTool.InvokableRun(ctx, string(argsBytes))
		if err != nil {
			glog.Errorf(ctx, "search_result error: %v", err)
		}
		glog.Debugf(ctx, "back_search_result: %s", result)
		state.BackgroundInvestigationResults = result
		return nil
	})
	return output, err
}

func bIRouter(ctx context.Context, input string, opts ...any) (output string, err error) {
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
		}()
		state.Goto = constants.AgentPlanner
		return nil
	})
	return output, nil
}

func NewBAgent[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()

	_ = cag.AddLambdaNode("search", compose.InvokableLambdaWithOption(search))
	_ = cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(bIRouter))

	_ = cag.AddEdge(compose.START, "search")
	_ = cag.AddEdge("search", "router")
	_ = cag.AddEdge("router", compose.END)
	return cag
}
