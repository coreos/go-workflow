package workflow_test

import (
	"testing"

	"github.com/colegleason/go-workflow"
)

func TestBasicWorkflow(t *testing.T) {
	var testVar bool

	w := workflow.New()
	w.AddStep(&workflow.Step{
		Label: "modify testVar",
		Run: func(c workflow.Context) error {
			testVar = true
			return nil
		},
	})

	err := w.Run()
	if err != nil {
		t.Error(err)
	}
	if testVar != true {
		t.Fail()
	}
}
