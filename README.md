go-workflow
-----------

[![Build Status](https://travis-ci.org/colegleason/go-workflow.svg?branch=master)](https://travis-ci.org/colegleason/go-workflow)

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
