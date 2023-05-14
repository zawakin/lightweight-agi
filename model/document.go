package model

import (
	"github.com/google/uuid"
)

type DocumentID uuid.UUID

func NewDocumentID() DocumentID {
	return DocumentID(uuid.New())
}

type Document struct {
	ID       DocumentID
	Text     string
	Metadata *DocumentMetadata
}

type Source string

type DocumentMetadata struct {
}

type DocumentChunkMetadata struct {
	*DocumentMetadata
	DocumentID DocumentID
}

type DocumentChunkID uuid.UUID

func NewDocumentChunkID() DocumentChunkID {
	return DocumentChunkID(uuid.New())
}

type DocumentChunk struct {
	ID         DocumentChunkID
	DocumentID DocumentID
	Text       string
	Metadata   DocumentChunkMetadata
	Embedding  []float32
}

type DocumentChunkWithScore struct {
	DocumentChunk
	Score float32
}

type DocumentWithChunks struct {
	Document
	Chunks []DocumentChunk
}

type DocumentMetadataFilter struct {
	DocumentID *DocumentID
	Source     *Source
	SourceID   *string
	// Author     *string
}

type Query struct {
	Query  string
	Filter *DocumentMetadataFilter
	TopK   int
}

type QueryWithEmbedding struct {
	Query
	Embedding Embedding
}

type QueryResult struct {
	Query   string
	Results []DocumentChunkWithScore
}
