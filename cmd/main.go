package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/datastore/providers/inmemory"
	"github.com/zawakin/lightweight-agi/examples/agi"
	"github.com/zawakin/lightweight-agi/examples/prompts"
	"github.com/zawakin/lightweight-agi/llmclient/provider"
	"github.com/zawakin/lightweight-agi/prompt"
)

var (
	// defaultOpenAICompletionModel = openai.GPT3Dot5Turbo
	defaultOpenAICompletionModel = openai.GPT4

	verbose = true
)

func init() {
	// load dotenv
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// disable log prefix
	log.SetFlags(0)
}

func main() {
	// read env
	openAIAPIKey := os.Getenv("OPENAI_API_KEY")
	if openAIAPIKey == "" {
		log.Fatal("OPENAI_API_KEY is required")
	}
	openAIOrgID := os.Getenv("OPENAI_ORG_ID")

	ctx := context.Background()

	// ==== Create OpenAI client ====
	openAIConfig := openai.DefaultConfig(openAIAPIKey)
	if openAIOrgID != "" {
		openAIConfig.OrgID = openAIOrgID
	}
	openaiClient := openai.NewClientWithConfig(openAIConfig)
	completionClient := provider.NewOpenAICompletionClient(openaiClient, defaultOpenAICompletionModel)
	embeddingClient := provider.NewOpenAIEmbeddingClient(openaiClient)

	// ==== Create data store provider ====
	dataStore := datastore.NewDataStore(inmemory.NewInMemoryDataStore(), embeddingClient)

	// === Create Prompt Runner ====
	runner := prompt.NewPromptRunner(completionClient, verbose)

	// ==== Run AGI agent ====
	agiAgent := agi.NewAGIAgent(runner, dataStore)

	// Define Global Objective of this AGI
	objective := prompts.Objective("Define the feature of good GPT prompt.")
	if err := agiAgent.RunAGIByObjective(ctx, objective); err != nil {
		log.Fatal(err)
	}
}
