package evaluationagent

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

func TestEvaluationTaskPrompt_Format(t *testing.T) {
	objective := model.Objective("Complete the project")
	task := model.Task{Name: "Design UI"}
	taskResult := model.TaskResult{ResultText: "Designed a user-friendly interface"}

	prompt := NewEvaluationTaskPrompt(objective, task, taskResult)

	want := `You are an AI who evaluates one task based on the following objective: Complete the project.
Evaluate the following task result with score(0-100) based on the task description: Design UI.

Template:
` + agentutil.TripleQuote + `
Result: bad result which does not achieve the objective
Score: 0
Reason: This is a bad result because...
` + agentutil.TripleQuote + `

` + agentutil.TripleQuote + `
Result: good result which achieves the objective completely
Score: 100
Reason: This is a good result because...
` + agentutil.TripleQuote + `

Result: Designed a user-friendly interface
`

	got := prompt.Format()
	if got != want {
		t.Errorf("Expected: %s, got: %s, diff %s", want, got, cmp.Diff(want, got))
	}
}
