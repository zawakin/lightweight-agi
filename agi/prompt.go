package agi

import "fmt"

var (
	tripleBacktick = "```"
)

type ObjectiveRefinementPrompt struct {
	Objective Objective
}

func NewObjectiveRefinementPrompt(objective Objective) ObjectiveRefinementPrompt {
	return ObjectiveRefinementPrompt{
		Objective: objective,
	}
}

func (o ObjectiveRefinementPrompt) Format() string {
	return fmt.Sprintf(`You are an AI tasked with refining the following objective.

Please provide a more specific objective that can be used to create a task.

Template:
%s
Objective: {original objective}
Refined objective: {refined objective}
%s

Objective: %s
Refined objective:`, tripleBacktick, tripleBacktick, o.Objective)
}

// MilestoneCreationPrompt is the prompt that is shown to the user when they are
// asked to create a new milestone.
type MilestoneCreationPrompt struct {
	Objective Objective
}

func NewMilestoneCreationPrompt(objective Objective) MilestoneCreationPrompt {
	return MilestoneCreationPrompt{
		Objective: objective,
	}
}

func (m MilestoneCreationPrompt) Format() string {
	return fmt.Sprintf(`You are an AI tasked with creating milestones for the following objective.

Please provide a list of milestones that can be used to achieve the objective.

Template:
%s
Objective: {original objective}
Milestones:
1. {milestone1}
2. {milestone2}
3. {milestone3}
...
%s

Objective: %s
Milestones:`, tripleBacktick, tripleBacktick, m.Objective)
}

// ExecutionTaskPrompt is the prompt that is shown to the user when they are
// asked to execute a task.
type ExecutionTaskPrompt struct {
	Objective       Objective
	CurrentTask     Task
	SolvedTasks     Tasks
	RelevantContext TaskRelevantContext
}

func NewExecutionTaskPrompt(objective Objective, currentTask Task, solvedTasks Tasks, relevantContext TaskRelevantContext) ExecutionTaskPrompt {
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

type EvaluationTaskPrompt struct {
	Objective  Objective
	Task       Task
	TaskResult TaskResult
}

func NewEvaluationTaskPrompt(objective Objective, task Task, taskResult TaskResult) EvaluationTaskPrompt {
	return EvaluationTaskPrompt{
		Objective:  objective,
		Task:       task,
		TaskResult: taskResult,
	}
}

func (e EvaluationTaskPrompt) Format() string {
	return fmt.Sprintf(`You are an AI who evaluates one task based on the following objective: %s.
Evaluate the following task result with score(0-100) based on the task description: %s.

Template:
%s
Result: bad result which does not achieve the objective
Score: 0
Reason: This is a bad result because...
%s

%s
Result: good result which achieves the objective completely
Score: 100
Reason: This is a good result because...
%s

Result: %s
`, e.Objective, e.Task.Name, tripleBacktick, tripleBacktick, tripleBacktick, tripleBacktick, e.TaskResult.ResultText)
}

// TaskCreationPrompt is the prompt that is shown to the user when they are
// asked to create a new task.
type TaskCreationPrompt struct {
	Objective      Objective
	Tasks          Tasks
	LastTaskResult *TaskResult
}

func NewTaskCreationPrompt(objective Objective, tasks Tasks, lastTaskResult *TaskResult) TaskCreationPrompt {
	return TaskCreationPrompt{
		Objective:      objective,
		Tasks:          tasks,
		LastTaskResult: lastTaskResult,
	}
}

func (t TaskCreationPrompt) Format() string {
	var prompt string
	if t.LastTaskResult == nil {
		prompt = fmt.Sprintf(`You are an task creation AI to create new tasks with the following objective: %s.
Create new tasks to be completed by the AI system.
Return the tasks as an array.

Template:
%s
1. {task1}
2. {task2}
3. {task3}
...
%s
		`, t.Objective, tripleBacktick, tripleBacktick)
	} else {
		prompt = fmt.Sprintf(`You are an task creation AI that uses the result of an execution agent to create new tasks with the following objective: %s.
The last completed task has the result: %s.
This result was based on this task description: %s.
These are incomplete tasks: %s.
Based on the result, create new tasks to be completed by the AI system that do not overlap with incomplete tasks.
Return the tasks as an array.

Template:
%s
1. {task1}
2. {task2}
3. {task3}
...
%s
`, t.Objective, t.LastTaskResult.ResultText, t.LastTaskResult.Task.Name, t.Tasks, tripleBacktick, tripleBacktick)

	}
	return prompt
}

// PriorizationPrompt is the prompt that is shown to the user when they are
// asked to prioritize a list of tasks.
type PriorizationPrompt struct {
	Objective Objective
	Tasks     Tasks
}

func NewPriorizationPrompt(objective Objective, tasks Tasks) PriorizationPrompt {
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
`, p.Tasks, p.Objective, tripleBacktick, tripleBacktick)
}
