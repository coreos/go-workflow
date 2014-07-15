package workflow

import (
	"fmt"
)

type FailureFunc func(err error, step *Step, context Context) error

func RetryFailure(tries int) FailureFunc {
	return func(err error, step *Step, context Context) error {
		var currError error
		for tries > 0 {
			currError := step.Run(context)
			if currError == nil {
				return nil
			}
			tries--
		}
		return currError
	}
}

func InteractiveFailure(err error, step *Step, context Context) error {
	fmt.Println(err)
	fmt.Printf("Step %s failed. Press ENTER to retry or C-c to quit.\n", step.Label)
	fmt.Scanln()
	nextErr := step.Run(context)
	if nextErr != nil {
		return InteractiveFailure(nextErr, step, context)
	}
	return nil
}
