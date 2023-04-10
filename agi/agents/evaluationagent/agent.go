package evaluationagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
)

var (
	maxTokensTaskCompletion = 500
)

// EvaluationAgent is responsible for evaluating the result of a task.
// It is also responsible for storing the context of the task.
type EvaluationAgent interface {
	EvaluateTask(ctx context.Context, objective model.Objective, task model.Task, taskResult model.TaskResult) (model.TaskEvaluation, error)
}

var _ EvaluationAgent = &EvaluationAgentImpl{}

type EvaluationAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewEvaluationAgent(completionClient llmclient.CompletionClient) *EvaluationAgentImpl {
	return &EvaluationAgentImpl{
		completionClient: completionClient,
	}
}

func (e *EvaluationAgentImpl) EvaluateTask(ctx context.Context, objective model.Objective, task model.Task, taskResult model.TaskResult) (model.TaskEvaluation, error) {
	prompt := EvaluationTaskPrompt{
		Objective:  objective,
		Task:       task,
		TaskResult: taskResult,
	}.Format()

	result, err := e.completionClient.Complete(ctx, prompt, maxTokensTaskCompletion)
	if err != nil {
		return model.TaskEvaluation{}, err
	}

	evaluation, err := ParseTaskEvaluationFromString(result)
	if err != nil {
		return model.TaskEvaluation{}, err
	}

	return evaluation, nil
}
