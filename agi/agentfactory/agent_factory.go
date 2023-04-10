package agentfactory

import (
	"github.com/zawakin/lightweight-agi/agi/agents/evaluationagent"
	"github.com/zawakin/lightweight-agi/agi/agents/executionagent"
	"github.com/zawakin/lightweight-agi/agi/agents/objectiverefinementagent"
	"github.com/zawakin/lightweight-agi/agi/agents/priorizationagent"
	"github.com/zawakin/lightweight-agi/agi/agents/taskcontextagent"
	"github.com/zawakin/lightweight-agi/agi/agents/taskcreationagent"
	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

// AgentFactory is a factory for creating AGI agents.
// It is used to inject dependencies into the AGI agents.
type AgentFactory interface {
	NewObjectiveRefinementAgent() objectiverefinementagent.ObjectiveRefinementAgent
	NewExecutionAgent() executionagent.ExecutionAgent
	NewEvaluationAgent() evaluationagent.EvaluationAgent
	NewTaskCreationAgent() taskcreationagent.TaskCreationAgent
	NewPriorizationAgent() priorizationagent.PriorizationAgent
	NewTaskContextAgent() taskcontextagent.TaskContextAgent
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

func (a *AgentFactoryImpl) NewObjectiveRefinementAgent() objectiverefinementagent.ObjectiveRefinementAgent {
	return objectiverefinementagent.NewObjectiveRefinementAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewExecutionAgent() executionagent.ExecutionAgent {
	return executionagent.NewExecutionAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewEvaluationAgent() evaluationagent.EvaluationAgent {
	return evaluationagent.NewEvaluationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewTaskCreationAgent() taskcreationagent.TaskCreationAgent {
	return taskcreationagent.NewTaskCreationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewPriorizationAgent() priorizationagent.PriorizationAgent {
	return priorizationagent.NewPriorizationAgent(a.completionClient)
}

func (a *AgentFactoryImpl) NewTaskContextAgent() taskcontextagent.TaskContextAgent {
	return taskcontextagent.NewTaskContextAgent(a.vectorStore, a.embeddingClient)
}
