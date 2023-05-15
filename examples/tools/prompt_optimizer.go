package tools

import (
	"context"

	"github.com/zawakin/lightweight-agi/examples/prompts"
	"github.com/zawakin/lightweight-agi/prompt"
)

func OptimizePrompt(ctx context.Context, runner *prompt.PromptRunner, original *prompt.Prompt, iterations int) (*prompt.Prompt, error) {
	p := original

	for i := 0; i < iterations; i++ {
		var result prompts.OptimizePromptOutput
		err := runner.Run(ctx, prompts.OptimizePromptPrompt, prompts.OptimizePromptInput{
			Original: p,
		}, &result)
		if err != nil {
			return nil, err
		}

		p = result.OptimizedPrompt
	}

	return p, nil
}
