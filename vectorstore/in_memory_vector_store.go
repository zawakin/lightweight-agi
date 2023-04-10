package vectorstore

import (
	"context"
	"errors"
	"log"
	"math"
	"sort"

	"github.com/google/uuid"
)

var _ VectorStore = (*InMemoryVectorStore)(nil)

type InMemoryVectorStore struct {
	vectors map[VectorRecordID]VectorRecord
}

func NewInMemoryVectorStore() *InMemoryVectorStore {
	return &InMemoryVectorStore{
		vectors: make(map[VectorRecordID]VectorRecord),
	}
}

func (i *InMemoryVectorStore) StoreVector(ctx context.Context, vector VectorRecord) error {
	if vector.ID == VectorRecordID(uuid.Nil) {
		return errors.New("vector id is empty")
	}
	if _, ok := i.vectors[vector.ID]; ok {
		log.Printf("vector already exists: %s\n", vector.ID)
		return errors.New("vector already exists")
	}
	i.vectors[vector.ID] = vector
	return nil
}

func (i *InMemoryVectorStore) GetVector(ctx context.Context, id VectorRecordID) (VectorRecord, error) {
	vector, ok := i.vectors[id]
	if !ok {
		return VectorRecord{}, errors.New("vector not found")
	}
	return vector, nil
}

func (i *InMemoryVectorStore) GetAllVector(ctx context.Context) ([]VectorRecord, error) {
	var result []VectorRecord
	for _, v := range i.vectors {
		result = append(result, v)
	}
	return result, nil
}

func (i *InMemoryVectorStore) QueryByVector(ctx context.Context, vector Vector, limit int) ([]VectorRecord, error) {
	var heap []VectorRecordWithSimilarity

	for _, v := range i.vectors {
		similarity := cosineSimilarity(v.Vector, vector)
		heap = append(heap, VectorRecordWithSimilarity{
			VectorRecord: v,
			Similarity:   similarity,
		})
	}

	sort.Slice(heap, func(i, j int) bool {
		return heap[i].Similarity > heap[j].Similarity
	})

	var result []VectorRecord
	for i := 0; i < limit; i++ {
		if i >= len(heap) {
			break
		}
		result = append(result, heap[i].VectorRecord)
	}
	return result, nil
}

func cosineSimilarity(a, b Vector) float32 {
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
