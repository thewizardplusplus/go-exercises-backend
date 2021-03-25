package registers

import (
	"context"
	"sync"

	"github.com/thewizardplusplus/go-exercises-backend/gateways/handlers"
)

// ConcurrentSolutionRegister ...
type ConcurrentSolutionRegister struct {
	ids               chan uint
	stoppingCtx       context.Context
	stoppingCtxCancel context.CancelFunc
	innerRegister     handlers.SolutionRegister
}

// NewConcurrentSolutionRegister ...
func NewConcurrentSolutionRegister(
	bufferSize int,
	innerRegister handlers.SolutionRegister,
) ConcurrentSolutionRegister {
	stoppingCtx, stoppingCtxCancel := context.WithCancel(context.Background())
	return ConcurrentSolutionRegister{
		ids:               make(chan uint, bufferSize),
		stoppingCtx:       stoppingCtx,
		stoppingCtxCancel: stoppingCtxCancel,
		innerRegister:     innerRegister,
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
	register.stoppingCtxCancel()
}

// Stop ...
func (register ConcurrentSolutionRegister) Stop() {
	close(register.ids)
	<-register.stoppingCtx.Done()
}
