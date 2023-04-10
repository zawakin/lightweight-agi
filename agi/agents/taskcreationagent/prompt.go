package taskcreationagent

import (
	"fmt"

	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

// MilestoneCreationPrompt is the prompt that is shown to the user when they are
// asked to create a new milestone.
type MilestoneCreationPrompt struct {
	Objective model.Objective
}

func NewMilestoneCreationPrompt(objective model.Objective) MilestoneCreationPrompt {
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
Milestones:`, agentutil.TripleQuote, agentutil.TripleQuote, m.Objective)
}

// TaskCreationPrompt is the prompt that is shown to the user when they are
// asked to create a new task.
type TaskCreationPrompt struct {
	Objective      model.Objective
	Tasks          model.Tasks
	LastTaskResult *model.TaskResult
}

func NewTaskCreationPrompt(objective model.Objective, tasks model.Tasks, lastTaskResult *model.TaskResult) TaskCreationPrompt {
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
		`, t.Objective, agentutil.TripleQuote, agentutil.TripleQuote)
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
`, t.Objective, t.LastTaskResult.ResultText, t.LastTaskResult.Task.Name, t.Tasks, agentutil.TripleQuote, agentutil.TripleQuote)

	}
	return prompt
}
