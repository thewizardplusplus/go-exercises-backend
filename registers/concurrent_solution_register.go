package registers

import (
	"context"
	"sync"

	"github.com/thewizardplusplus/go-exercises-backend/entities"
)

// ConcurrentSolutionRegister ...
type ConcurrentSolutionRegister struct {
	innerRegister entities.SolutionRegister

	startMode            *startModeHolder
	stoppingCtx          context.Context
	stoppingCtxCanceller context.CancelFunc
	ids                  chan uint
}

// NewConcurrentSolutionRegister ...
func NewConcurrentSolutionRegister(
	bufferSize int,
	innerRegister entities.SolutionRegister,
) ConcurrentSolutionRegister {
	startMode := &startModeHolder{}
	stoppingCtx, stoppingCtxCanceller := context.WithCancel(context.Background())
	return ConcurrentSolutionRegister{
		innerRegister: innerRegister,

		startMode:            startMode,
		stoppingCtx:          stoppingCtx,
		stoppingCtxCanceller: stoppingCtxCanceller,
		ids:                  make(chan uint, bufferSize),
	}
}

// RegisterSolution ...
func (register ConcurrentSolutionRegister) RegisterSolution(id uint) {
	register.ids <- id
}

// Start ...
func (register ConcurrentSolutionRegister) Start() {
	register.basicRun(started, func() {
		for id := range register.ids {
			register.innerRegister.RegisterSolution(id)
		}
	})
}

// StartConcurrently ...
func (register ConcurrentSolutionRegister) StartConcurrently(concurrency int) {
	register.basicRun(startedConcurrently, func() {
		var waitGroup sync.WaitGroup
		waitGroup.Add(concurrency)

		for threadID := 0; threadID < concurrency; threadID++ {
			go func() {
				defer waitGroup.Done()

				register.Start()
			}()
		}

		waitGroup.Wait()
	})
}

// Stop ...
func (register ConcurrentSolutionRegister) Stop() {
	close(register.ids)
	<-register.stoppingCtx.Done()
}

func (register ConcurrentSolutionRegister) basicRun(
	mode startMode,
	runHandler func(),
) {
	register.startMode.setStartModeOnce(mode)

	runHandler()

	if register.startMode.getStartMode() == mode {
		register.stoppingCtxCanceller()
	}
}
