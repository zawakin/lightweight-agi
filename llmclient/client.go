package llmclient

import (
	"context"

	"github.com/zawakin/lightweight-agi/vectorstore"
)

type EmbeddingClient interface {
	EmbedText(ctx context.Context, text string) (vectorstore.Vector, error)
}

// CompletionClient is an interface that provides a method to complete text.
type CompletionClient interface {
	Complete(ctx context.Context, text string, maxTokens int) (string, error)
}
