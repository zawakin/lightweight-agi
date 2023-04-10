package objectiverefinementagent

import (
	"fmt"

	"github.com/zawakin/lightweight-agi/agi/agentutil"
	"github.com/zawakin/lightweight-agi/agi/model"
)

type ObjectiveRefinementPrompt struct {
	Objective model.Objective
}

func NewObjectiveRefinementPrompt(objective model.Objective) ObjectiveRefinementPrompt {
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
Refined objective:`, agentutil.TripleQuote, agentutil.TripleQuote, o.Objective)
}
