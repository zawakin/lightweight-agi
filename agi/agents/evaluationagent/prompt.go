package evaluationagent

import (
	"fmt"

	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

type EvaluationTaskPrompt struct {
	Objective  model.Objective
	Task       model.Task
	TaskResult model.TaskResult
}

func NewEvaluationTaskPrompt(objective model.Objective, task model.Task, taskResult model.TaskResult) EvaluationTaskPrompt {
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
`, e.Objective, e.Task.Name, agentutil.TripleQuote, agentutil.TripleQuote, agentutil.TripleQuote, agentutil.TripleQuote, e.TaskResult.ResultText)
}
