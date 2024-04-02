package question

import (
	"fmt"
)

type Processor struct {
	pointer   int
	questions []*Question
}

func (p *Processor) GetQuestion() (*Question, error) {
	if p.pointer < 0 || p.pointer >= len(p.questions) {
		return nil, fmt.Errorf("питання не існує")
	}

	return p.questions[p.pointer], nil
}

func (p *Processor) Increment() {
	p.pointer++
}

func (p *Processor) IsLast() bool {
	return p.pointer+1 == len(p.questions)
}

func NewProcessor(questions []*Question) *Processor {
	return &Processor{questions: questions}
}
