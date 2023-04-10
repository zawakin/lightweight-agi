package taskcreationagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
)

var (
	maxTokensTasksList = 200
)

// TaskCreationAgent is responsible for creating new tasks.
// It is also responsible for storing the context of the task.
type TaskCreationAgent interface {
	CreateMilestones(ctx context.Context, objective model.Objective) (model.Milestones, error)
	CreateTasks(ctx context.Context, objective model.Objective, tasks model.Tasks, taskResult *model.TaskResult) (model.Tasks, error)
}

var _ TaskCreationAgent = &TaskCreationAgentImpl{}

type TaskCreationAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewTaskCreationAgent(completionClient llmclient.CompletionClient) *TaskCreationAgentImpl {
	return &TaskCreationAgentImpl{
		completionClient: completionClient,
	}
}

func (t *TaskCreationAgentImpl) CreateTasks(ctx context.Context, ojbective model.Objective, tasks model.Tasks, lastTaskResult *model.TaskResult) (model.Tasks, error) {
	prompt := NewTaskCreationPrompt(ojbective, tasks, lastTaskResult).Format()
	result, err := t.completionClient.Complete(ctx, prompt, maxTokensTasksList)
	if err != nil {
		return model.Tasks{}, err
	}

	return ParseTasksFromString(result)
}

func (m *TaskCreationAgentImpl) CreateMilestones(ctx context.Context, objective model.Objective) (model.Milestones, error) {
	prompt := NewMilestoneCreationPrompt(objective).Format()
	result, err := m.completionClient.Complete(ctx, prompt, maxTokensTasksList)
	if err != nil {
		return nil, err
	}

	milestones, err := ParseMilestones(result)
	if err != nil {
		return nil, err
	}

	return milestones, nil
}
