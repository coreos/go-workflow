package workflow_test

import (
	"testing"

	"github.com/colegleason/go-workflow"
)

func TestBasicWorkflow(t *testing.T) {
	var testVar bool

	step := &workflow.Step{
		Label: "modify testVar",
		Run: func(c workflow.Context) error {
			testVar = true
			return nil
		},
	}

	w := workflow.New()
	w.Start = step
	err := w.Run()
	if err != nil {
		t.Error(err)
	}
	if testVar != true {
		t.Fail()
	}
}

func TestDependancyWorkflow(t *testing.T) {
	var testVars [4]bool

	one := &workflow.Step{
		Label: "modify testVar 1",
		Run: func(c workflow.Context) error {
			testVars[1] = true
			return nil
		},
	}

	three := &workflow.Step{
		Label: "modify testVar 3",
		Run: func(c workflow.Context) error {
			testVars[3] = true
			return nil
		},
	}

	two := &workflow.Step{
		Label:     "modify testVar 2",
		DependsOn: []*workflow.Step{three},
		Run: func(c workflow.Context) error {
			if !testVars[3] {
				t.Fail()
			}
			testVars[2] = true
			return nil
		},
	}

	base := &workflow.Step{
		Label:     "modify testVar 0",
		DependsOn: []*workflow.Step{one, two},
		Run: func(c workflow.Context) error {
			if !testVars[1] || !testVars[2] {
				t.Fail()
			}
			testVars[0] = true
			return nil
		},
	}

	w := workflow.New()
	w.Start = base
	err := w.Run()
	if err != nil {
		t.Error(err)
	}
	if testVars[0] != true {
		t.Fail()
	}
}
