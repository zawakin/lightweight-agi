package datastore

import (
	"context"

	"github.com/zawakin/lightweight-agi/llmclient"
	"github.com/zawakin/lightweight-agi/model"
	"github.com/zawakin/lightweight-agi/services"
)

type DataStore struct {
	provider        DataStoreProvider
	embeddingClient llmclient.EmbeddingClient
}

type DataStoreProvider interface {
	Upsert(ctx context.Context, chunks map[model.DocumentID][]model.DocumentChunk, chunkTokenSize *int) ([]model.DocumentID, error)
	Query(ctx context.Context, queries []model.QueryWithEmbedding) ([]model.QueryResult, error)
}

func NewDataStore(provider DataStoreProvider, embeddingClient llmclient.EmbeddingClient) *DataStore {
	return &DataStore{
		provider:        provider,
		embeddingClient: embeddingClient,
	}
}

func (d *DataStore) Upsert(ctx context.Context, documents []model.Document, chunkTokenSize *int) ([]model.DocumentID, error) {
	chunks := make(map[model.DocumentID][]model.DocumentChunk)
	for _, doc := range documents {
		documentChunks, err := services.CreateDocumentChunks(ctx, d.embeddingClient, &doc, *chunkTokenSize)
		if err != nil {
			return nil, err
		}
		chunks[doc.ID] = documentChunks
	}

	ids, err := d.provider.Upsert(ctx, chunks, chunkTokenSize)
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (d *DataStore) Query(ctx context.Context, queries []model.Query) ([]model.QueryResult, error) {
	queryTexts := make([]string, 0, len(queries))
	for _, query := range queries {
		queryTexts = append(queryTexts, query.Query)
	}

	queryEmbeddings, err := d.embeddingClient.EmbedTexts(ctx, queryTexts)
	if err != nil {
		return nil, err
	}

	queryWithEmbeddings := make([]model.QueryWithEmbedding, len(queries))
	for i, query := range queries {
		queryWithEmbeddings[i] = model.QueryWithEmbedding{
			Query:     query,
			Embedding: queryEmbeddings[i],
		}
	}

	return d.provider.Query(ctx, queryWithEmbeddings)
}
