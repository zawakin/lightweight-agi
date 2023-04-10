package provider

import (
	"context"
	"errors"

	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/vectorstore"
)

var _ llmclient.EmbeddingClient = (*OpenAIEmbeddingClient)(nil)

type OpenAIEmbeddingClient struct {
	openAIClient *openai.Client
}

func NewOpenAIEmbeddingClient(openAIClient *openai.Client) *OpenAIEmbeddingClient {
	return &OpenAIEmbeddingClient{
		openAIClient: openAIClient,
	}
}

func (e *OpenAIEmbeddingClient) EmbedText(ctx context.Context, text string) (vectorstore.Vector, error) {
	result, err := e.openAIClient.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: []string{text},
		Model: openai.AdaEmbeddingV2,
	})
	if err != nil {
		return nil, err
	}
	if len(result.Data) == 0 {
		return nil, errors.New("no embeddings returned")
	}
	return result.Data[0].Embedding, nil
}

var _ llmclient.CompletionClient = (*OpenAICompletionClient)(nil)

type OpenAICompletionClient struct {
	openAIClient *openai.Client
	model        string
}

func NewOpenAICompletionClient(openAIClient *openai.Client, model string) *OpenAICompletionClient {
	return &OpenAICompletionClient{
		openAIClient: openAIClient,
		model:        model,
	}
}

func (c *OpenAICompletionClient) Complete(ctx context.Context, text string, maxTokens int) (string, error) {
	result, err := c.openAIClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: c.model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "user",
				Content: text,
			},
		},
		MaxTokens:   maxTokens,
		Temperature: 0.7,
	})
	if err != nil {
		return "", err
	}
	if len(result.Choices) == 0 {
		return "", errors.New("no choices returned")
	}
	return result.Choices[0].Message.Content, nil
}
