package priorizationagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
)

var (
	maxTokensTasksList = 200
)

// PriorizationAgent is responsible for prioritizing tasks.
type PriorizationAgent interface {
	PrioritizeTasks(ctx context.Context, objective model.Objective, tasks model.Tasks) (model.Tasks, error)
}

type PriorizationAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewPriorizationAgent(completionClient llmclient.CompletionClient) *PriorizationAgentImpl {
	return &PriorizationAgentImpl{
		completionClient: completionClient,
	}
}

func (p *PriorizationAgentImpl) PrioritizeTasks(ctx context.Context, objective model.Objective, tasks model.Tasks) (model.Tasks, error) {
	prompt := NewPriorizationPrompt(objective, tasks).Format()
	result, err := p.completionClient.Complete(ctx, prompt, maxTokensTasksList)
	if err != nil {
		return model.Tasks{}, err
	}

	newTasks, err := ParseTasksFromString(result)
	if err != nil {
		return model.Tasks{}, err
	}

	return newTasks, nil
}
