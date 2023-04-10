package agi

import (
	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

type AgentFactory interface {
	NewObjectiveRefinementAgent() ObjectiveRefinementAgent
	NewExecutionAgent() ExecutionAgent
	NewEvaluationAgent() EvaluationAgent
	NewTaskCreationAgent() TaskCreationAgent
	NewPriorizationAgent() PriorizationAgent
	NewTaskContextAgent() TaskContextAgent
}

type AgentFactoryImpl struct {
	vectorStore      vectorstore.VectorStore
	completionClient llmclient.CompletionClient
	embeddingClient  llmclient.EmbeddingClient
}

func NewAgentFactory(
	vectorStore vectorstore.VectorStore,
	completionClient llmclient.CompletionClient,
	embeddingClient llmclient.EmbeddingClient,
) AgentFactory {
	return &AgentFactoryImpl{
		vectorStore:      vectorStore,
		completionClient: completionClient,
		embeddingClient:  embeddingClient,
	}
}

func (a *AgentFactoryImpl) NewObjectiveRefinementAgent() ObjectiveRefinementAgent {
	return NewObjectiveRefinementAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewExecutionAgent() ExecutionAgent {
	return NewExecutionAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewEvaluationAgent() EvaluationAgent {
	return NewEvaluationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewTaskCreationAgent() TaskCreationAgent {
	return NewTaskCreationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewPriorizationAgent() PriorizationAgent {
	return NewPriorizationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewTaskContextAgent() TaskContextAgent {
	return NewTaskContextAgent(a.vectorStore, a.embeddingClient)
}
