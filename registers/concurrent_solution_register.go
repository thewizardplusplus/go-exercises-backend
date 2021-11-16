package registers

import (
	"context"

	"github.com/thewizardplusplus/go-exercises-backend/entities"
	syncutils "github.com/thewizardplusplus/go-sync-utils"
)

type solutionRegisterWrapper struct {
	solutionRegister entities.SolutionRegister
}

func (wrapper solutionRegisterWrapper) Handle(
	ctx context.Context,
	data interface{},
) {
	wrapper.solutionRegister.RegisterSolution(data.(uint))
}

// ConcurrentSolutionRegister ...
type ConcurrentSolutionRegister struct {
	// do not use embedding to hide the Handle() method
	concurrentHandler syncutils.ConcurrentHandler
}

// NewConcurrentSolutionRegister ...
func NewConcurrentSolutionRegister(
	bufferSize int,
	innerSolutionRegister entities.SolutionRegister,
) ConcurrentSolutionRegister {
	return ConcurrentSolutionRegister{
		concurrentHandler: syncutils.NewConcurrentHandler(
			bufferSize,
			solutionRegisterWrapper{
				solutionRegister: innerSolutionRegister,
			},
		),
	}
}

// RegisterSolution ...
func (register ConcurrentSolutionRegister) RegisterSolution(id uint) {
	register.concurrentHandler.Handle(id)
}

// Start ...
func (register ConcurrentSolutionRegister) Start() {
	register.concurrentHandler.Start(context.Background())
}

// StartConcurrently ...
func (register ConcurrentSolutionRegister) StartConcurrently(
	concurrencyFactor int,
) {
	register.concurrentHandler.
		StartConcurrently(context.Background(), concurrencyFactor)
}

// Stop ...
func (register ConcurrentSolutionRegister) Stop() {
	register.concurrentHandler.Stop()
}
