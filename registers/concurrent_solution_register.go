package registers

import (
	"sync"

	"github.com/thewizardplusplus/go-exercises-backend/gateways/handlers"
)

// ConcurrentSolutionRegister ...
type ConcurrentSolutionRegister struct {
	ids           chan uint
	innerRegister handlers.SolutionRegister
}

// NewConcurrentSolutionRegister ...
func NewConcurrentSolutionRegister(
	bufferSize int,
	innerRegister handlers.SolutionRegister,
) ConcurrentSolutionRegister {
	return ConcurrentSolutionRegister{
		ids:           make(chan uint, bufferSize),
		innerRegister: innerRegister,
	}
}

// RegisterSolution ...
func (register ConcurrentSolutionRegister) RegisterSolution(id uint) {
	register.ids <- id
}

// Start ...
func (register ConcurrentSolutionRegister) Start() {
	for id := range register.ids {
		register.innerRegister.RegisterSolution(id)
	}
}

// StartConcurrently ...
func (register ConcurrentSolutionRegister) StartConcurrently(concurrency int) {
	var waitGroup sync.WaitGroup
	waitGroup.Add(concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			defer waitGroup.Done()

			register.Start()
		}()
	}

	waitGroup.Wait()
}

// Stop ...
func (register ConcurrentSolutionRegister) Stop() {
	close(register.ids)
}
