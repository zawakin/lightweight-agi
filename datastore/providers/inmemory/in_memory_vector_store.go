package inmemory

import (
	"context"
	"math"
	"sort"

	"github.com/zawakin/lightweight-agi/datastore"
	"github.com/zawakin/lightweight-agi/model"
)

var _ datastore.DataStoreProvider = (*InMemoryDataStore)(nil)

type InMemoryDataStore struct {
	data map[model.DocumentChunkID]model.DocumentChunk
}

func NewInMemoryDataStore() *InMemoryDataStore {
	return &InMemoryDataStore{
		data: make(map[model.DocumentChunkID]model.DocumentChunk),
	}
}

func (s *InMemoryDataStore) Upsert(ctx context.Context, chunks map[model.DocumentID][]model.DocumentChunk, chunkTokenSize *int) ([]model.DocumentID, error) {
	var result []model.DocumentID
	for docID, v := range chunks {
		for _, chunk := range v {
			s.data[chunk.ID] = chunk
		}
		result = append(result, docID)
	}
	return result, nil
}

func (s *InMemoryDataStore) Query(ctx context.Context, queries []model.QueryWithEmbedding) ([]model.QueryResult, error) {
	var result []model.QueryResult
	for _, query := range queries {
		r, err := s.query(ctx, query)
		if err != nil {
			return nil, err
		}
		result = append(result, *r)
	}
	return result, nil
}

func (s *InMemoryDataStore) query(ctx context.Context, query model.QueryWithEmbedding) (*model.QueryResult, error) {
	var heap []model.DocumentChunkWithScore
	topK := query.TopK

	for _, chunk := range s.data {
		score := cosineSimilarity(query.Embedding.Embedding, chunk.Embedding)
		heap = append(heap, model.DocumentChunkWithScore{
			DocumentChunk: chunk,
			Score:         score,
		})
	}

	sort.Slice(heap, func(i, j int) bool {
		return heap[i].Score > heap[j].Score
	})
	if len(heap) > topK {
		heap = heap[:topK]
	}

	var result []model.DocumentChunkWithScore
	result = append(result, heap...)

	return &model.QueryResult{
		Query:   query.Query.Query,
		Results: result,
	}, nil

}

func cosineSimilarity(a, b []float32) float32 {
	var dotProduct, aMagnitude, bMagnitude float32

	for i := range a {
		dotProduct += a[i] * b[i]
		aMagnitude += a[i] * a[i]
		bMagnitude += b[i] * b[i]
	}

	aMagnitude = float32(math.Sqrt(float64(aMagnitude)))
	bMagnitude = float32(math.Sqrt(float64(bMagnitude)))

	return dotProduct / (aMagnitude * bMagnitude)
}
