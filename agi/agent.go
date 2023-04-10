package agi

import (
	"context"
)

// ObjectiveRefinementAgent is responsible for refining the objective.
type ObjectiveRefinementAgent interface {
	RefineObjective(ctx context.Context, objective Objective) (Objective, error)
}

// ExecutionAgent is responsible for executing a task.
// It is also responsible for storing the context of the task.
type ExecutionAgent interface {
	ExecuteTask(ctx context.Context, objective Objective, task Task, relevantContext TaskRelevantContext) (TaskResult, error)
}

// EvaluationAgent is responsible for evaluating the result of a task.
// It is also responsible for storing the context of the task.
type EvaluationAgent interface {
	EvaluateTask(ctx context.Context, objective Objective, task Task, taskResult TaskResult) (TaskEvaluation, error)
}

// TaskCreationAgent is responsible for creating new tasks.
// It is also responsible for storing the context of the task.
type TaskCreationAgent interface {
	CreateMilestones(ctx context.Context, objective Objective) (Milestones, error)
	CreateTasks(ctx context.Context, objective Objective, tasks Tasks, taskResult *TaskResult) (Tasks, error)
}

// PriorizationAgent is responsible for prioritizing tasks.
type PriorizationAgent interface {
	PrioritizeTasks(ctx context.Context, objective Objective, tasks Tasks) (Tasks, error)
}

// TaskContextAgent is responsible for storing the context of a task.
// It is also responsible for retrieving the context of a task.
type TaskContextAgent interface {
	FindRelevantContext(ctx context.Context, task Task) (TaskRelevantContext, error)
	StoreContext(ctx context.Context, taskContext TaskContext) error

	// DebugDumpTaskContext returns the entire context of the agent.
	// This is used for debugging purposes.
	DebugDumpTaskContext(ctx context.Context) ([]TaskContext, error)
}
