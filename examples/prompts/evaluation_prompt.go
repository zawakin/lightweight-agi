package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type EvaluationTaskInput struct {
	Objective  Objective  `json:"objective"`
	Task       Task       `json:"task"`
	TaskResult TaskResult `json:"task_result"`
}

type EvaluationTaskOutput struct {
	Score  int    `json:"score"`
	Reason string `json:"reason"`
}

var (
	EvaluationTasksPrompt = &prompt.Prompt{
		Name:        "evaluation tasks",
		Description: `Evaluate the following task result with score(0-100) based on the task description.`,
		Template: &prompt.Example{
			Input: &EvaluationTaskInput{
				Objective:  Objective("objective"),
				Task:       Task{Name: "task"},
				TaskResult: TaskResult{ResultText: "result text"},
			},
			Output: &EvaluationTaskOutput{
				Score:  50,
				Reason: "reason",
			},
		},
		Examples: prompt.Examples{
			{
				Input: &EvaluationTaskInput{
					Objective:  Objective("objective"),
					Task:       Task{Name: "task-1"},
					TaskResult: TaskResult{ResultText: "bad result which does not achieve the objective"},
				},
				Output: &EvaluationTaskOutput{
					Score:  0,
					Reason: "This is a bad result because...",
				},
			},
			{
				Input: &EvaluationTaskInput{
					Objective:  Objective("objective"),
					Task:       Task{Name: "task"},
					TaskResult: TaskResult{ResultText: "good result which achieves the objective completely"},
				},
				Output: &EvaluationTaskOutput{
					Score:  100,
					Reason: "This is a good result because...",
				},
			},
		},
	}
)
