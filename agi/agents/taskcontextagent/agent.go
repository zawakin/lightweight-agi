package taskcontextagent

import (
	"context"

	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

// TaskContextAgent is responsible for storing the context of a task.
// It is also responsible for retrieving the context of a task.
type TaskContextAgent interface {
	FindRelevantContext(ctx context.Context, task model.Task) (model.TaskRelevantContext, error)
	StoreContext(ctx context.Context, taskContext model.TaskContext) error

	// DebugDumpTaskContext returns the entire context of the agent.
	// This is used for debugging purposes.
	DebugDumpTaskContext(ctx context.Context) ([]model.TaskContext, error)
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

func (t *TaskContextAgentImpl) FindRelevantContext(ctx context.Context, task model.Task) (model.TaskRelevantContext, error) {
	embedding, err := t.embeddingClient.EmbedText(ctx, task.Name)
	if err != nil {
		return model.TaskRelevantContext{}, err
	}

	vectors, err := t.vectorStore.QueryByVector(ctx, embedding, 5)
	if err != nil {
		return model.TaskRelevantContext{}, err
	}

	var relevantContext []model.TaskContext
	for _, vector := range vectors {
		taskContext, err := model.GetTaskContentFromMetadata(vector.Metadata)
		if err != nil {
			return model.TaskRelevantContext{}, err
		}

		relevantContext = append(relevantContext, taskContext)
	}

	return model.TaskRelevantContext{
		BaseTask:        task,
		RelevantContext: relevantContext,
	}, nil
}

func (t *TaskContextAgentImpl) DebugDumpTaskContext(ctx context.Context) ([]model.TaskContext, error) {
	vectors, err := t.vectorStore.GetAllVector(ctx)
	if err != nil {
		return nil, err
	}

	contexts := make([]model.TaskContext, 0, len(vectors))
	for _, vector := range vectors {
		taskContext, err := model.GetTaskContentFromMetadata(vector.Metadata)
		if err != nil {
			return nil, err
		}

		contexts = append(contexts, taskContext)
	}

	return contexts, nil
}

func (t *TaskContextAgentImpl) StoreContext(ctx context.Context, taskContext model.TaskContext) error {
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
