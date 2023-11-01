package service

import (
	"context"
	"lang-bot/internal/word"
	"time"

	"github.com/google/uuid"
)

type WordService struct {
	wordRepository word.Repository
}

func NewWordService(wordRepository word.Repository) WordService {
	return WordService{
		wordRepository: wordRepository,
	}

}

func (s WordService) AddWord(ctx context.Context, text string) (word.Word, error) {
	newWord := word.Word{
		ID:        uuid.New(),
		Text:      text,
		CreatedAt: time.Now(),
	}

	err := s.wordRepository.AddWord(ctx, newWord)
	if err != nil {
		return word.Word{}, err
	}

	return newWord, nil
}
