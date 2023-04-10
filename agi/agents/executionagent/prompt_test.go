package executionagent

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zawakin/lightweight-agi/agi/model"
)

func TestExecutionTaskPrompt_Format(t *testing.T) {
	objective := model.Objective("Complete the project")
	currentTask := model.Task{Name: "Design UI"}
	solvedTasks := model.Tasks{
		{Name: "Define requirements"},
		{Name: "Create wireframes"},
	}
	relevantContext := model.TaskRelevantContext{
		BaseTask: currentTask,
		// "Consider the user experience",
		RelevantContext: []model.TaskContext{
			{Task: model.Task{Name: "relevant_task_name_1"}, Text: "relevant_task_text_1"},
			{Task: model.Task{Name: "relevant_task_name_2"}, Text: "relevant_task_text_2"},
		},
	}

	prompt := NewExecutionTaskPrompt(objective, currentTask, solvedTasks, relevantContext)

	want := `Objective: Complete the project
Previously completed tasks: ["Define requirements","Create wireframes"]
Current task: Design UI
Relevant context: relevant_task_text_1,relevant_task_text_2

Execute the current task, considering the relevant context and previous tasks.

Response:`

	got := prompt.Format()
	if got != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, got, cmp.Diff(want, got))
	}
}
