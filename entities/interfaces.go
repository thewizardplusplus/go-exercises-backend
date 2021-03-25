package entities

// TaskGetter ...
type TaskGetter interface {
	GetTask(id uint) (Task, error)
}

// SolutionRegister ...
type SolutionRegister interface {
	RegisterSolution(id uint)
}
