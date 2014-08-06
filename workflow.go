package workflow

import (
	"fmt"
)

type Context interface{}

type Workflow struct {
	Start     *Step
	OnFailure FailureFunc
	Context   Context

	queue   []*Step
	inQueue map[*Step]bool
}

func New() *Workflow {
	w := &Workflow{}
	w.inQueue = make(map[*Step]bool)
	return w
}

func (w *Workflow) Run() error {
	w.loadQueue(w.Start)
	for _, step := range w.queue {
		fmt.Printf("Running step: %s ", step.Label)
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

func (w *Workflow) loadQueue(s *Step) {
	if s == nil {
		return
	}

	for _, step := range s.DependsOn {
		w.loadQueue(step)
	}

	if !w.inQueue[s] {
		w.inQueue[s] = true
		w.queue = append(w.queue, s)
	}
	return
}
