/*
 * @Author: morehao morehao@qq.com
 * @Date: 2025-08-11 10:25:29
 * @LastEditors: morehao morehao@qq.com
 * @LastEditTime: 2025-09-13 22:05:07
 * @FilePath: /einodeer/engine/investigator.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package engine

import (
	"context"
	"encoding/json"

	"github.com/cloudwego/eino/compose"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func search(ctx context.Context, name string, opts ...any) (output string, err error) {
	// 使用新的工具系统获取搜索工具
	searchTool := infra.DefaultToolManager.GetToolByNameSuffix("search")

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
