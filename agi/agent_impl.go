package agi

import (
	"context"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

var (
	maxTokensTasksList           = 200
	maxTokensTaskCompletion      = 500
	maxTokensObjectiveRefinement = 200
)

type ObjectiveRefinementAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewObjectiveRefinementAgent(completionClient llmclient.CompletionClient) *ObjectiveRefinementAgentImpl {
	return &ObjectiveRefinementAgentImpl{
		completionClient: completionClient,
	}
}

func (o *ObjectiveRefinementAgentImpl) RefineObjective(ctx context.Context, objective Objective) (Objective, error) {
	prompt := NewObjectiveRefinementPrompt(objective).Format()
	result, err := o.completionClient.Complete(ctx, prompt, maxTokensObjectiveRefinement)
	if err != nil {
		return "", err
	}

	return Objective(result), nil
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

func (e *ExecutionAgentImpl) ExecuteTask(ctx context.Context, objective Objective, task Task, relevantContext TaskRelevantContext) (TaskResult, error) {
	prompt := ExecutionTaskPrompt{
		Objective:       objective,
		CurrentTask:     task,
		RelevantContext: relevantContext,
	}.Format()

	result, err := e.completionClient.Complete(ctx, prompt, maxTokensTaskCompletion)
	if err != nil {
		return TaskResult{}, err
	}

	return TaskResult{
		Task:       task,
		ResultText: result,
	}, nil
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

func (e *EvaluationAgentImpl) EvaluateTask(ctx context.Context, objective Objective, task Task, taskResult TaskResult) (TaskEvaluation, error) {
	prompt := EvaluationTaskPrompt{
		Objective:  objective,
		Task:       task,
		TaskResult: taskResult,
	}.Format()

	result, err := e.completionClient.Complete(ctx, prompt, maxTokensTaskCompletion)
	if err != nil {
		return TaskEvaluation{}, err
	}

	evaluation, err := ParseTaskEvaluationFromString(result)
	if err != nil {
		return TaskEvaluation{}, err
	}

	return evaluation, nil
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

func (t *TaskCreationAgentImpl) CreateTasks(ctx context.Context, ojbective Objective, tasks Tasks, lastTaskResult *TaskResult) (Tasks, error) {
	prompt := NewTaskCreationPrompt(ojbective, tasks, lastTaskResult).Format()
	result, err := t.completionClient.Complete(ctx, prompt, maxTokensTasksList)
	if err != nil {
		return Tasks{}, err
	}

	return ParseTasksFromString(result)
}

func (m *TaskCreationAgentImpl) CreateMilestones(ctx context.Context, objective Objective) (Milestones, error) {
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

type PriorizationAgentImpl struct {
	completionClient llmclient.CompletionClient
}

func NewPriorizationAgent(completionClient llmclient.CompletionClient) *PriorizationAgentImpl {
	return &PriorizationAgentImpl{
		completionClient: completionClient,
	}
}

func (p *PriorizationAgentImpl) PrioritizeTasks(ctx context.Context, objective Objective, tasks Tasks) (Tasks, error) {
	prompt := NewPriorizationPrompt(objective, tasks).Format()
	result, err := p.completionClient.Complete(ctx, prompt, maxTokensTasksList)
	if err != nil {
		return Tasks{}, err
	}

	newTasks, err := ParseTasksFromString(result)
	if err != nil {
		return Tasks{}, err
	}

	return newTasks, nil
}

var _ TaskContextAgent = &TaskContextAgentImpl{}

type TaskContextAgentImpl struct {
	embeddingClient llmclient.EmbeddingClient
	vectorStore     vectorstore.VectorStore
}

func NewTaskContextAgent(vectorStore vectorstore.VectorStore, embeddingClient llmclient.EmbeddingClient) *TaskContextAgentImpl {
	return &TaskContextAgentImpl{
		vectorStore:     vectorStore,
		embeddingClient: embeddingClient,
	}
}

func (t *TaskContextAgentImpl) FindRelevantContext(ctx context.Context, task Task) (TaskRelevantContext, error) {
	embedding, err := t.embeddingClient.EmbedText(ctx, task.Name)
	if err != nil {
		return TaskRelevantContext{}, err
	}

	vectors, err := t.vectorStore.QueryByVector(ctx, embedding, 5)
	if err != nil {
		return TaskRelevantContext{}, err
	}

	var relevantContext []TaskContext
	for _, vector := range vectors {
		taskContext, err := GetTaskContentFromMetadata(vector.Metadata)
		if err != nil {
			return TaskRelevantContext{}, err
		}

		relevantContext = append(relevantContext, taskContext)
	}

	return TaskRelevantContext{
		BaseTask:        task,
		RelevantContext: relevantContext,
	}, nil
}

func (t *TaskContextAgentImpl) DebugDumpTaskContext(ctx context.Context) ([]TaskContext, error) {
	vectors, err := t.vectorStore.GetAllVector(ctx)
	if err != nil {
		return nil, err
	}

	contexts := make([]TaskContext, 0, len(vectors))
	for _, vector := range vectors {
		taskContext, err := GetTaskContentFromMetadata(vector.Metadata)
		if err != nil {
			return nil, err
		}

		contexts = append(contexts, taskContext)
	}

	return contexts, nil
}

func (t *TaskContextAgentImpl) StoreContext(ctx context.Context, taskContext TaskContext) error {
	md := map[string]interface{}{
		"id":      taskContext.Task.ID.String(),
		"task":    taskContext.Task.Name,
		"content": taskContext.Text,
	}

	embedding, err := t.embeddingClient.EmbedText(ctx, taskContext.Text)
	if err != nil {
		return err
	}

	err = t.vectorStore.StoreVector(ctx, vectorstore.VectorRecord{
		ID:       vectorstore.VectorRecordID(taskContext.Task.ID),
		Vector:   embedding,
		Metadata: md,
	})
	if err != nil {
		return err
	}

	return nil
}
