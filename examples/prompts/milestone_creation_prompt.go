package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type MilestoneCreationInput struct {
	Objective Objective `json:"objective"`
}

type MilestoneCreationOutput struct {
	Milestones Milestones `json:"milestones"`
}

var (
	MilestoneCreationPrompt = &prompt.Prompt{
		Name: "milestone creation",
		Description: `You are an AI tasked with creating a milestone for the following objective.

Please provide a milestone that can be used to achieve the objective.`,
		Template: &prompt.Example{
			Input: &MilestoneCreationInput{
				Objective: Objective("original objective"),
			},
			Output: &MilestoneCreationOutput{
				Milestones: Milestones{
					{Objective: Objective("milestone1"), Name: "milestone1"},
					{Objective: Objective("milestone2"), Name: "milestone2"},
					{Objective: Objective("milestone3"), Name: "milestone3"},
				},
			},
		},
	}
)

type TaskCreationInput struct {
	Objective      Objective   `json:"objective"`
	QueuedTasks    Tasks       `json:"queued_tasks"`
	LastTaskResult *TaskResult `json:"last_task_result"`
}

type TaskCreationOutput struct {
	Tasks Tasks `json:"tasks"`
}

var (
	TaskCreationPrompt = &prompt.Prompt{
		Name:        "task creation",
		Description: `Create new tasks to be completed in order to achieve the objective if necessary.`,
		Template: &prompt.Example{
			Input: &TaskCreationInput{
				Objective: Objective("original objective"),
				QueuedTasks: Tasks{
					Task{
						Name: "task-10",
					},
					Task{
						Name: "task-11",
					},
				},
				LastTaskResult: &TaskResult{
					Task: Task{
						Name: "task-9",
					},
					ResultText: "task-9 result",
				},
			},
			Output: &TaskCreationOutput{
				Tasks: Tasks{
					Task{
						Name: "task-10",
					},
					Task{
						Name: "task-11",
					},
					Task{
						Name: "task-12",
					},
				},
			},
		},
	}
)
