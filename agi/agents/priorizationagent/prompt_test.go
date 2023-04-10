package priorizationagent

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

func TestPriorizationPrompt_Format(t *testing.T) {
	objective := model.Objective("Complete the project")
	tasks := model.Tasks{
		{Name: "Define requirements"},
		{Name: "Create wireframes"},
		{Name: "Design UI"},
	}

	prompt := NewPriorizationPrompt(objective, tasks)

	want := `You are a task prioritization AI responsible for organizing the following tasks in a higher-priority order: ["Define requirements","Create wireframes","Design UI"].
Consider the ultimate objective of your team: Complete the project.

To prioritize these tasks, please follow the steps below:

1. Determine the importance of each task based on the ultimate objective.
2. Consider any dependencies between tasks or any external constraints (e.g., deadlines, resources) that may impact the order of execution.
3. Reorder the tasks accordingly, with the most important and urgent tasks at the top.

Do not remove any tasks. Return the tasks as an array in higher-priority order, using the following format:

Template:
` + agentutil.TripleQuote + `
1. {task1}
2. {task2}
3. {task3}
...
` + agentutil.TripleQuote + `
`

	got := prompt.Format()
	if got != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, prompt.Format(), cmp.Diff(want, got))
	}
}
