package transfer

import "testing"

func Test_BuildTree(t *testing.T) {
	var nodes = []NodeItem{
		{Id: 1, Pid: 0, Label: "1"},
		{Id: 2, Pid: 1, Label: "1-2"},
		{Id: 3, Pid: 1, Label: "1-3"},
		{Id: 4, Pid: 3, Label: "1-3-4"},
		{Id: 5, Pid: 0, Label: "5"},
		{Id: 6, Pid: 5, Label: "5-6"},
		{Id: 7, Pid: 6, Label: "5-6-7"},
		{Id: 8, Pid: 6, Label: "5-6-8"},
	}
	BuildTree(nodes)
}
