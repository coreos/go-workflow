package workflow

type StepFunc func(context interface{}) error

type Step struct {
	Label string
	Run   StepFunc
}
