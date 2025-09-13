package model

import (
	"fmt"
	"testing"
)

func TestState(t *testing.T) {
	state := State{
		Messages:          nil,
		Goto:              "",
		PlanIterations:    0,
		MaxPlanIterations: 0,
		MaxStepNum:        0,
		CurrentPlan: &Plan{
			Thought: "asdas",
			Steps: []Step{
				{
					Title: "asdas",
				},
			},
		},
		Locale: "",
		//Server:                         nil,
		InterruptFeedback:              "",
		AutoAcceptedPlan:               false,
		EnableBackgroundInvestigation:  false,
		BackgroundInvestigationResults: "",
	}
	bt, err := state.MarshalJSON()
	//bt, err := json.Marshal(state)
	if err != nil {
		fmt.Println("编码失败,错误原因: ", err)
		return
	}
	fmt.Println(string(bt))
}
