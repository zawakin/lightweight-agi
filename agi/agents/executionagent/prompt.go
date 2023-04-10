package executionagent

import (
	"fmt"

	"github.com/zawakin/lightweight-agi/agi/model"
)

// ExecutionTaskPrompt is the prompt that is shown to the user when they are
// asked to execute a task.
type ExecutionTaskPrompt struct {
	Objective       model.Objective
	CurrentTask     model.Task
	SolvedTasks     model.Tasks
	RelevantContext model.TaskRelevantContext
}

func NewExecutionTaskPrompt(objective model.Objective, currentTask model.Task, solvedTasks model.Tasks, relevantContext model.TaskRelevantContext) ExecutionTaskPrompt {
	return ExecutionTaskPrompt{
		Objective:       objective,
		CurrentTask:     currentTask,
		SolvedTasks:     solvedTasks,
		RelevantContext: relevantContext,
	}
}

func (e ExecutionTaskPrompt) Format() string {
	return fmt.Sprintf(`Objective: %s
Previously completed tasks: %s
Current task: %s
Relevant context: %s

Execute the current task, considering the relevant context and previous tasks.

Response:`, e.Objective, e.SolvedTasks.String(), e.CurrentTask.Name, e.RelevantContext)
}
