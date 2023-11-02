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

func (s WordLocalStorage) GetWords(ctx context.Context) ([]word.Word, error) {
	words := make([]word.Word, 0)

	for _, w := range s.words {
		words = append(words, w)
	}

	return words, nil
}
