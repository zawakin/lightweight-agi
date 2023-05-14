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

	// ==== Run AGI agent ====
	agiAgent := agi.NewAGIAgent(
		completionClient,
		embeddingClient,
		dataStore,
	)

	// Define Global Objective of this AGI
	objective := prompts.Objective("I want to learn how to play chess.")

	if err := agiAgent.RunAGIByObjective(ctx, objective); err != nil {
		log.Fatal(err)
	}

	// 	p := prompt.NewSimplePrompt(
	// 		"Fix JSON grammar",
	// 		`Fix JSON grammar.
	// Don't output any code block like markdown.
	// Just output raw JSON.Format tightly removing whitespaces between keys and values.`,
	// 		"<original JSON>",
	// 		"<fixed JSON>",
	// 	)

	p := prompts.EvaluationTasksPrompt
	// input := prompts.EvaluationTaskInput{
	// 	Objective: prompts.Objective("I want to learn how to play chess."),
	// 	Task:      prompts.Task{},
	// }

	runner := prompt.NewPromptRunner(completionClient)
	var resp prompts.RefinePromptOutput

	for i := 0; i < 5; i++ {
		err := runner.Run(ctx, prompts.RefinePromptPrompt, prompts.RefinePromptInput{
			Original: p,
		}, &resp)
		if err != nil {
			log.Println(err)
			continue
		}
		p = resp.RefinedPrompt
	}
}
