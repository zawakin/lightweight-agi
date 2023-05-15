package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

// PrioritizationPrompt is the prompt that is shown to the user when they are
// asked to prioritize a list of tasks.
type PrioritizationInput struct {
	Objective Objective `json:"objective"`
	Tasks     Tasks     `json:"tasks"`
}

type PriorizationOutput struct {
	Tasks []Task `json:"tasks"`
}

var (
	PrioritizationPrompt = &prompt.Prompt{
		Name: "task prioritization",
		Description: `You are a task prioritization AI responsible for organizing the following tasks in a higher-priority order.

To prioritize these tasks, please follow the steps below:

1. Determine the importance of each task based on the ultimate objective.
2. Consider any dependencies between tasks or any external constraints (e.g., deadlines, resources) that may impact the order of execution.
3. Reorder the tasks accordingly, with the most important and urgent tasks at the top.

Do not remove any tasks. Return the tasks as an array in higher-priority order.
`,
		Template: &prompt.Example{
			Input: &PrioritizationInput{
				Objective: Objective("Objective"),
				Tasks: Tasks{
					{Name: "Task 1"},
					{Name: "Task 2"},
					{Name: "Task 3"},
				},
			},
			Output: &PriorizationOutput{
				Tasks: []Task{
					{Name: "Task 2"},
					{Name: "Task 1"},
					{Name: "Task 3"},
				},
			},
		},
	}
)
