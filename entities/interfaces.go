package entities

// TaskGetter ...
type TaskGetter interface {
	GetTask(userID uint, taskID uint) (Task, error)
}

// SolutionGetter ...
type SolutionGetter interface {
	GetSolution(id uint) (Solution, error)
}

// SolutionRegister ...
type SolutionRegister interface {
	RegisterSolution(id uint)
}
