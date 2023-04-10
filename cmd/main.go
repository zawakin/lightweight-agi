package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/agi"
	"github.com/zawakin/lightweight-agi/agi/agentfactory"
	"github.com/zawakin/lightweight-agi/agi/model"
	"github.com/zawakin/lightweight-agi/provider"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

var (
	defaultOpenAICompletionModel = openai.GPT3Dot5Turbo
)

func main() {
	// load dotenv
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// read env OPENAI_API_KEY
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")

	// Define Global Objective of this AGI
	objective := model.Objective("I want to learn how to play chess.")

	// disable log prefix
	log.SetFlags(0)

	ctx := context.Background()

	// ==== Create OpenAI client ====
	openaiClient := openai.NewClient(openAIAPIKey)
	completionClient := provider.NewOpenAICompletionClient(openaiClient, defaultOpenAICompletionModel)
	embeddingClient := provider.NewOpenAIEmbeddingClient(openaiClient)

	// ==== Create vector store ====
	vectorStore := vectorstore.NewInMemoryVectorStore()

	// ==== Create AGI agent factory ====
	agentFactory := agentfactory.NewAgentFactory(vectorStore, completionClient, embeddingClient)

	// ==== Run AGI agent ====
	agiAgent := agi.NewAGIAgent(
		agentFactory.NewObjectiveRefinementAgent(),
		agentFactory.NewExecutionAgent(),
		agentFactory.NewEvaluationAgent(),
		agentFactory.NewTaskCreationAgent(),
		agentFactory.NewPriorizationAgent(),
		agentFactory.NewTaskContextAgent(),
	)
	if err := agiAgent.RunAGIByObjective(ctx, objective); err != nil {
		log.Fatal(err)
	}
}
