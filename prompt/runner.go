package prompt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/zawakin/lightweight-agi/llmclient"
)

type PromptRunner struct {
	llmClient llmclient.CompletionClient
}

func NewPromptRunner(llmClient llmclient.CompletionClient) *PromptRunner {
	return &PromptRunner{
		llmClient: llmClient,
	}
}

func (a *PromptRunner) Run(ctx context.Context, prompter Prompter, input Input, out any) error {
	prompt, err := prompter.Format(input)
	if err != nil {
		return err
	}

	log.Println("--------------------------------")
	log.Printf("\033[33mPrompt:\n%s\033[0m\n", prompt)

	result, err := a.llmClient.Complete(ctx, prompt, 1000)
	if err != nil {
		return err
	}

	fmt.Printf("\033[32mResult:\n%s\033[0m\n", result)
	fmt.Println("--------------------------------")

	err = json.Unmarshal([]byte(result), out)
	if err != nil {
		return err
	}

	return nil
}
