package workflow_test

import (
	"errors"
	"testing"

	"github.com/colegleason/go-workflow"
)

func TestFailureFunc(t *testing.T) {
	var testVar bool

	w := workflow.New()
	w.OnFailure = func(err error, step *workflow.Step, context workflow.Context) error {
		testVar = true
		return nil
	}
	w.AddStep(&workflow.Step{
		Label: "fail workflow",
		Run: func(c workflow.Context) error {
			return errors.New("generic error")
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
