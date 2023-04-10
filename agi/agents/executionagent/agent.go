package executionagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
)

var (
	maxTokensTaskCompletion = 500
)

// ExecutionAgent is responsible for executing a task.
// It is also responsible for storing the context of the task.
type ExecutionAgent interface {
	ExecuteTask(ctx context.Context, objective model.Objective, task model.Task, relevantContext model.TaskRelevantContext) (model.TaskResult, error)
}

var _ ExecutionAgent = &ExecutionAgentImpl{}

type ExecutionAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewExecutionAgent(completionClient llmclient.CompletionClient) *ExecutionAgentImpl {
	return &ExecutionAgentImpl{
		completionClient: completionClient,
	}
}

func (e *ExecutionAgentImpl) ExecuteTask(ctx context.Context, objective model.Objective, task model.Task, relevantContext model.TaskRelevantContext) (model.TaskResult, error) {
	prompt := ExecutionTaskPrompt{
		Objective:       objective,
		CurrentTask:     task,
		RelevantContext: relevantContext,
	}.Format()

	result, err := e.completionClient.Complete(ctx, prompt, maxTokensTaskCompletion)
	if err != nil {
		return model.TaskResult{}, err
	}

	return model.TaskResult{
		Task:       task,
		ResultText: result,
	}, nil
}
