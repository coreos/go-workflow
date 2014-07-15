package workflow

import (
	"errors"
	"fmt"
)

type Workflow struct {
	StepOrder []string
	Steps     map[string]*Step
	OnFailure FailureFunc
	Context   interface{}
}

func New() *Workflow {
	w := &Workflow{}
	w.Steps = make(map[string]*Step)
	return w
}

func (w *Workflow) AddStep(s *Step) error {
	if w.Steps[s.Label] != nil {
		return fmt.Errorf("%s already exists in workflow", s.Label)
	}
	w.Steps[s.Label] = s
	w.StepOrder = append(w.StepOrder, s.Label)
	return nil
}

func (w *Workflow) AddSteps(steps []*Step) error {
	oldStepOrder := w.StepOrder
	oldStepList := w.Steps
	for _, step := range steps {
		err := w.AddStep(step)
		if err != nil {
			w.StepOrder = oldStepOrder
			w.Steps = oldStepList
			return err
		}
	}
	return nil
}

func (w *Workflow) Run() error {
	if len(w.StepOrder) < 1 {
		return errors.New("Workflow contains no steps.")
	}
	return w.RunFrom(w.StepOrder[0])
}

func (w *Workflow) RunFrom(stepName string) error {
	start := w.getStepIndex(stepName)
	if start == -1 {
		return fmt.Errorf("Step %s not found in workflow.", stepName)
	}
	for _, stepName := range w.StepOrder[start:] {
		step := w.Steps[stepName]
		if step == nil {
			return fmt.Errorf("Step %s does not exist in workflow.", stepName)
		}
		fmt.Printf("Running step: %s ", stepName)
		if err := step.Run(w.Context); err != nil {
			if err := w.OnFailure(err, step, w.Context); err != nil {
				fmt.Println("FAILED")
				return err
			}
		}
		fmt.Println("COMPLETE")
	}
	return nil

}

func (w *Workflow) getStepIndex(stepName string) int {
	for i, step := range w.StepOrder {
		if stepName == step {
			return i
		}
	}
	return -1
}
