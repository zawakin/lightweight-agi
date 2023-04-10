package objectiverefinementagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
)

var (
	maxTokensObjectiveRefinement = 200
)

// ObjectiveRefinementAgent is responsible for refining the objective.
type ObjectiveRefinementAgent interface {
	RefineObjective(ctx context.Context, objective model.Objective) (model.Objective, error)
}

type ObjectiveRefinementAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewObjectiveRefinementAgent(completionClient llmclient.CompletionClient) *ObjectiveRefinementAgentImpl {
	return &ObjectiveRefinementAgentImpl{
		completionClient: completionClient,
	}
}

func (o *ObjectiveRefinementAgentImpl) RefineObjective(ctx context.Context, objective model.Objective) (model.Objective, error) {
	prompt := NewObjectiveRefinementPrompt(objective).Format()
	result, err := o.completionClient.Complete(ctx, prompt, maxTokensObjectiveRefinement)
	if err != nil {
		return "", err
	}

	return model.Objective(result), nil
}
