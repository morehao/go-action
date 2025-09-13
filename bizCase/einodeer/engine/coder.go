package engine

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
	"github.com/cloudwego/eino/schema"
	"github.com/morehao/go-action/bizCase/einodeer/constants"
	"github.com/morehao/go-action/bizCase/einodeer/infra"
	"github.com/morehao/go-action/bizCase/einodeer/model"
	"github.com/morehao/golib/glog"
)

func loadCoderMsg(ctx context.Context, name string, opts ...any) (output []*schema.Message, err error) {
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

		var curStep *model.Step
		for i := range state.CurrentPlan.Steps {
			if state.CurrentPlan.Steps[i].ExecutionRes == nil {
				curStep = &state.CurrentPlan.Steps[i]
				break
			}
		}

		if curStep == nil {
			panic("no step found")
		}

		msg := []*schema.Message{}
		msg = append(msg,
			schema.UserMessage(fmt.Sprintf("#Task\n\n##title\n\n %v \n\n##description\n\n %v \n\n##locale\n\n %v", curStep.Title, curStep.Description, state.Locale)),
		)
		variables := map[string]any{
			"locale":              state.Locale,
			"max_step_num":        state.MaxStepNum,
			"max_plan_iterations": state.MaxPlanIterations,
			"CURRENT_TIME":        time.Now().Format("2006-01-02 15:04:05"),
			"user_input":          msg,
		}
		output, err = promptTemp.Format(ctx, variables)
		return err
	})
	return output, err
}

func routerCoder(ctx context.Context, input *schema.Message, opts ...any) (output string, err error) {
	last := input
	err = compose.ProcessState[*model.State](ctx, func(_ context.Context, state *model.State) error {
		defer func() {
			output = state.Goto
		}()
		for i, step := range state.CurrentPlan.Steps {
			if step.ExecutionRes == nil {
				str := strings.Clone(last.Content)
				state.CurrentPlan.Steps[i].ExecutionRes = &str
				break
			}
		}
		glog.Infof(ctx, "coder_end: %s", glog.ToJsonString(state.CurrentPlan))
		state.Goto = constants.AgentResearchTeam
		return nil
	})
	return output, nil
}

func modifyCoderfunc(ctx context.Context, input []*schema.Message) []*schema.Message {
	sum := 0
	maxLimit := 50000
	for i := range input {
		if input[i] == nil {
			glog.Warnf(ctx, "modify_inputfunc_nil: %s", glog.ToJsonString(input[i]))
			continue
		}
		l := len(input[i].Content)
		if l > maxLimit {
			glog.Warnf(ctx, "modify_inputfunc_clip: %s", glog.ToJsonString(input[i]))
			input[i].Content = input[i].Content[l-maxLimit:]
		}
		sum += len(input[i].Content)
	}
	glog.Infof(ctx, "modify_inputfunc sum: %d, input: %s", sum, glog.ToJsonString(input))
	return input
}

func NewCoder[I, O any](ctx context.Context) *compose.Graph[I, O] {
	cag := compose.NewGraph[I, O]()

	// 使用新的工具系统获取工具
	researchTools := []tool.BaseTool{}
	// 获取所有工具并转换为BaseTool列表
	tools := infra.DefaultToolManager.GetAllTools()
	for name, t := range tools {
		// 只添加Golang相关工具
		if strings.HasPrefix(name, "golang") {
			if baseTool, ok := t.(tool.BaseTool); ok {
				researchTools = append(researchTools, baseTool)
			}
		}
	}
	glog.Debugf(ctx, "coder_end coder_tools: %s", glog.ToJsonString(researchTools))

	agent, err := react.NewAgent(ctx, &react.AgentConfig{
		MaxStep:               40,
		ToolCallingModel:      infra.ChatModel,
		ToolsConfig:           compose.ToolsNodeConfig{Tools: researchTools},
		MessageModifier:       modifyCoderfunc,
		StreamToolCallChecker: toolCallChecker,
	})
	if err != nil {
		glog.Errorf(ctx, "coder_end agent: %v", err)
		return nil
	}

	agentLambda, err := compose.AnyLambda(agent.Generate, agent.Stream, nil, nil)
	if err != nil {
		glog.Errorf(ctx, "coder_end agent_lambda: %v", err)
		return nil
	}

	if err := cag.AddLambdaNode("load", compose.InvokableLambdaWithOption(loadCoderMsg)); err != nil {
		glog.Errorf(ctx, "coder_end load: %v", err)
		return nil
	}
	if err := cag.AddLambdaNode("agent", agentLambda); err != nil {
		glog.Errorf(ctx, "coder_end agent: %v", err)
		return nil
	}
	if err := cag.AddLambdaNode("router", compose.InvokableLambdaWithOption(routerCoder)); err != nil {
		glog.Errorf(ctx, "coder_end router: %v", err)
		return nil
	}

	if err := cag.AddEdge(compose.START, "load"); err != nil {
		glog.Errorf(ctx, "coder_end load: %v", err)
		return nil
	}
	if err := cag.AddEdge("load", "agent"); err != nil {
		glog.Errorf(ctx, "coder_end agent: %v", err)
		return nil
	}
	if err := cag.AddEdge("agent", "router"); err != nil {
		glog.Errorf(ctx, "coder_end router: %v", err)
		return nil
	}
	if err := cag.AddEdge("router", compose.END); err != nil {
		glog.Errorf(ctx, "coder_end end: %v", err)
		return nil
	}
	return cag
}
