package booking

import (
	"06_interface/booking/step"
	"fmt"
)

type Processor struct {
	steps       []step.Step
	stepPointer int
}

func (p *Processor) GetStep() (step.Step, error) {
	if p.stepPointer < 0 || p.stepPointer >= len(p.steps) {
		return nil, fmt.Errorf("крок не існую")
	}

	return p.steps[p.stepPointer], nil
}

func (p *Processor) IncrementStep() {
	p.stepPointer++
}

func (p *Processor) IsLastStep() bool {
	return p.stepPointer+1 == len(p.steps)
}

func BuildBookProcessor() *Processor {
	steps := step.BuildBookingSteps()

	return &Processor{
		steps: steps,
	}
}
