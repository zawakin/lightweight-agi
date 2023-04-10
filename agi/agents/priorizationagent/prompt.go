package priorizationagent

import (
	"fmt"

	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

// PriorizationPrompt is the prompt that is shown to the user when they are
// asked to prioritize a list of tasks.
type PriorizationPrompt struct {
	Objective model.Objective
	Tasks     model.Tasks
}

func NewPriorizationPrompt(objective model.Objective, tasks model.Tasks) PriorizationPrompt {
	return PriorizationPrompt{
		Objective: objective,
		Tasks:     tasks,
	}
}

func (p PriorizationPrompt) Format() string {
	return fmt.Sprintf(`You are a task prioritization AI responsible for organizing the following tasks in a higher-priority order: %s.
Consider the ultimate objective of your team: %s.

To prioritize these tasks, please follow the steps below:

1. Determine the importance of each task based on the ultimate objective.
2. Consider any dependencies between tasks or any external constraints (e.g., deadlines, resources) that may impact the order of execution.
3. Reorder the tasks accordingly, with the most important and urgent tasks at the top.

Do not remove any tasks. Return the tasks as an array in higher-priority order, using the following format:

Template:
%s
1. {task1}
2. {task2}
3. {task3}
...
%s
`, p.Tasks, p.Objective, agentutil.TripleQuote, agentutil.TripleQuote)
}
