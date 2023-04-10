package taskcreationagent

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

func TestTaskCreationPrompt_Format(t *testing.T) {
	objective := model.Objective("Complete the project")
	tasks := model.Tasks{
		{Name: "Define requirements"},
		{Name: "Create wireframes"},
	}

	prompt := NewTaskCreationPrompt(objective, tasks, nil)

	want := `You are an task creation AI to create new tasks with the following objective: Complete the project.
Create new tasks to be completed by the AI system.
Return the tasks as an array.

Template:
` + agentutil.TripleQuote + `
1. {task1}
2. {task2}
3. {task3}
...
` + agentutil.TripleQuote + `
		`

	got := prompt.Format()
	if prompt.Format() != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, prompt.Format(), cmp.Diff(want, got))
	}
}
