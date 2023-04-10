package objectiverefinementagent

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

func TestObjectiveRefinementPrompt_Format(t *testing.T) {
	objective := model.Objective("Complete the project")

	prompt := NewObjectiveRefinementPrompt(objective)

	want := `You are an AI tasked with refining the following objective.

Please provide a more specific objective that can be used to create a task.

Template:
` + agentutil.TripleQuote + `
Objective: {original objective}
Refined objective: {refined objective}
` + agentutil.TripleQuote + `

Objective: Complete the project
Refined objective:`
	got := prompt.Format()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("prompt.Format() mismatch (-want +got):\n%s", diff)
	}
}
