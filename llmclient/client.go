package llmclient

import (
	"context"

	"github.com/zawakin/lightweight-agi/model"
)

type EmbeddingClient interface {
	EmbedText(ctx context.Context, text string) (*model.Embedding, error)
	EmbedTexts(ctx context.Context, text []string) ([]model.Embedding, error)
}

// CompletionClient is an interface that provides a method to complete text.
type CompletionClient interface {
	Complete(ctx context.Context, text string, maxTokens int) (string, error)
}

type ChatCompletionClient interface {
	Complete(ctx context.Context, messages ChatMessages, opt CompletionOption) (string, error)
}

type Role string

const (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
	RoleSystem    Role = "system"
)

func (r Role) String() string {
	return string(r)
}

type ChatMessage struct {
	Role    Role
	Content string
}

type ChatMessages []ChatMessage

type CompletionOption struct {
	MaxTokens int
}
