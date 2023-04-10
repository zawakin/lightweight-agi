package agi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestObjectiveRefinementPrompt_Format(t *testing.T) {
	objective := Objective("Complete the project")

	prompt := NewObjectiveRefinementPrompt(objective)

	want := `You are an AI tasked with refining the following objective.

Please provide a more specific objective that can be used to create a task.

Template:
` + tripleBacktick + `
Objective: {original objective}
Refined objective: {refined objective}
` + tripleBacktick + `

Objective: Complete the project
Refined objective:`
	got := prompt.Format()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("prompt.Format() mismatch (-want +got):\n%s", diff)
	}
}

func TestExecutionTaskPrompt_Format(t *testing.T) {
	objective := Objective("Complete the project")
	currentTask := Task{Name: "Design UI"}
	solvedTasks := Tasks{
		{Name: "Define requirements"},
		{Name: "Create wireframes"},
	}
	relevantContext := TaskRelevantContext{
		BaseTask: currentTask,
		// "Consider the user experience",
		RelevantContext: []TaskContext{
			{Task: Task{Name: "relevant_task_name_1"}, Text: "relevant_task_text_1"},
			{Task: Task{Name: "relevant_task_name_2"}, Text: "relevant_task_text_2"},
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

func TestEvaluationTaskPrompt_Format(t *testing.T) {
	objective := Objective("Complete the project")
	task := Task{Name: "Design UI"}
	taskResult := TaskResult{ResultText: "Designed a user-friendly interface"}

	prompt := NewEvaluationTaskPrompt(objective, task, taskResult)

	want := `You are an AI who evaluates one task based on the following objective: Complete the project.
Evaluate the following task result with score(0-100) based on the task description: Design UI.

Template:
` + tripleBacktick + `
Result: bad result which does not achieve the objective
Score: 0
Reason: This is a bad result because...
` + tripleBacktick + `

` + tripleBacktick + `
Result: good result which achieves the objective completely
Score: 100
Reason: This is a good result because...
` + tripleBacktick + `

Result: Designed a user-friendly interface
`

	got := prompt.Format()
	if got != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, got, cmp.Diff(want, got))
	}
}

func TestTaskCreationPrompt_Format(t *testing.T) {
	objective := Objective("Complete the project")
	tasks := Tasks{
		{Name: "Define requirements"},
		{Name: "Create wireframes"},
	}

	prompt := NewTaskCreationPrompt(objective, tasks, nil)

	want := `You are an task creation AI to create new tasks with the following objective: Complete the project.
Create new tasks to be completed by the AI system.
Return the tasks as an array.

Template:
` + tripleBacktick + `
1. {task1}
2. {task2}
3. {task3}
...
` + tripleBacktick + `
		`

	got := prompt.Format()
	if prompt.Format() != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, prompt.Format(), cmp.Diff(want, got))
	}
}

func TestPriorizationPrompt_Format(t *testing.T) {
	objective := Objective("Complete the project")
	tasks := Tasks{
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
` + tripleBacktick + `
1. {task1}
2. {task2}
3. {task3}
...
` + tripleBacktick + `
`

	got := prompt.Format()
	if got != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, prompt.Format(), cmp.Diff(want, got))
	}
}
