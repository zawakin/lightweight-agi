package vectorstore

import (
	"context"

	"github.com/google/uuid"
)

type VectorRecordID uuid.UUID

func MakeVectorRecordID() VectorRecordID {
	return VectorRecordID(uuid.New())
}

func (id VectorRecordID) String() string {
	return uuid.UUID(id).String()
}

func (id VectorRecordID) MarshalText() ([]byte, error) {
	return []byte(uuid.UUID(id).String()), nil
}

func ParseVectorRecordIDFromString(s string) (VectorRecordID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return VectorRecordID{}, err
	}
	return VectorRecordID(id), nil
}

// VectorRecord is a struct that contains the base vector record information.
// It is used to identify the vector record and to provide a name for the vector record.
type VectorRecord struct {
	ID       VectorRecordID
	Vector   Vector
	Metadata map[string]interface{}
}

type Vector []float32

// VectorRecordWithSimilarity is a struct that contains the vector record information
// and the similarity to the given vector.
type VectorRecordWithSimilarity struct {
	VectorRecord
	Similarity float32 // cosine similarity
}

// VectorStore is an interface that provides methods to store and retrieve vectors.
type VectorStore interface {
	// StoreVector stores the given vector record.
	// If the vector record already exists, it will cause an error.
	StoreVector(ctx context.Context, vector VectorRecord) error

	// GetVector returns the vector record with the given ID.
	GetVector(ctx context.Context, id VectorRecordID) (VectorRecord, error)

	// GetAllVector returns all vector records.
	GetAllVector(ctx context.Context) ([]VectorRecord, error)

	// QueryByVector returns the vector records that are most similar to the given vector.
	// The limit parameter specifies the maximum number of results to return.
	QueryByVector(ctx context.Context, vector Vector, limit int) ([]VectorRecord, error)
}
