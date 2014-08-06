package workflow

import (
	"fmt"
	"io"
	"os"
)

var (
	InputFile io.Reader = os.Stdin
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
	fmt.Fprintln(os.Stderr, err)
	fmt.Fprintf(os.Stdout, "Step %s failed. Press r to retry, s to skip, or C-c to quit.\n", step.Label)
	for {
		var action string
		_, err := fmt.Fscanln(InputFile, &action)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(1)
		}
		switch action {
		case "r":
			nextErr := step.Run(context)
			if nextErr != nil {
				return InteractiveFailure(nextErr, step, context)
			} else {
				return nil
			}
		case "s":
			return nil
		}
		fmt.Fprintf(os.Stdout, "Invalid action '%s'. Valid actions=[r,s]\n", action)
	}
}
