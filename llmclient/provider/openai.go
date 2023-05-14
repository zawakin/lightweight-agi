package provider

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/sashabaranov/go-openai"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
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

func (e *OpenAIEmbeddingClient) EmbedText(ctx context.Context, text string) (*model.Embedding, error) {
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
	return &model.Embedding{
		Embedding: result.Data[0].Embedding,
	}, nil
}

func (e *OpenAIEmbeddingClient) EmbedTexts(ctx context.Context, texts []string) ([]model.Embedding, error) {
	result, err := e.openAIClient.CreateEmbeddings(ctx, openai.EmbeddingRequest{
		Input: texts,
		Model: openai.AdaEmbeddingV2,
	})
	if err != nil {
		return nil, err
	}
	embeddings := make([]model.Embedding, len(result.Data))
	for i, data := range result.Data {
		embeddings[i] = model.Embedding{
			Embedding: data.Embedding,
		}
	}
	return embeddings, nil
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
	result, err := c.complete(ctx, text, maxTokens)
	if err != nil {
		log.Println("retrying completion after 1 second", err)
		time.Sleep(1 * time.Second)

		result, err = c.complete(ctx, text, maxTokens)
		if err != nil {
			return "", err
		}
	}
	return result, nil
}

func (c *OpenAICompletionClient) complete(ctx context.Context, text string, maxTokens int) (string, error) {
	resp, err := c.openAIClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
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
	if len(resp.Choices) == 0 {
		return "", errors.New("no choices returned")
	}
	result := resp.Choices[0].Message.Content
	return result, nil
}
