package localcache

import (
	"context"
	"lang-bot/internal/word"

	"github.com/google/uuid"
)

type WordLocalStorage struct {
	words map[uuid.UUID]word.Word
}

func NewWordLocalStorage() WordLocalStorage {
	return WordLocalStorage{
		words: make(map[uuid.UUID]word.Word),
	}
}

func (s WordLocalStorage) AddWord(ctx context.Context, w word.Word) error {
	s.words[w.ID] = w

	return nil
}
