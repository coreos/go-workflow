package workflow

type StepFunc func(context Context) error

type Step struct {
	Label     string
	Run       StepFunc
	DependsOn []*Step
}
