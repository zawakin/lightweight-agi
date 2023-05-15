package prompts

import (
	"github.com/zawakin/lightweight-agi/prompt"
)

type ObjectiveRefinementInput struct {
	Objective Objective `json:"objective"`
}

type ObjectiveRefinementOutput struct {
	RefinedObjective Objective `json:"refined_objective"`
}

var (
	ObjectRefinementPrompt = &prompt.Prompt{
		Name:        "refinement of objective",
		Description: `Refine the objective to a more specific objective that can be used.`,
		Template: &prompt.Example{
			Input: &ObjectiveRefinementInput{
				Objective: Objective("original objective"),
			},
			Output: &ObjectiveRefinementOutput{
				RefinedObjective: Objective("refined objective"),
			},
		},
	}
)
