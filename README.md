go-workflow
-----------

[![deprecated](http://badges.github.io/stability-badges/dist/deprecated.svg)](http://github.com/badges/stability-badges)


Simple control flow library to setup a series of steps to execute.

Example
-------

```
w := workflow.New()
w.OnFailure = workflow.InteractiveFailure
steps := []*workflow.Step{
	&workflow.Step{
		Label: "one",
		Run:   stepOne,
	},
	&workflow.Step{
		Label: "two",
		Run:   stepTwo,
		},
	}
}
w.AddSteps(steps)
w.Run()
```
