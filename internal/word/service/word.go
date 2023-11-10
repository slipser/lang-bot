package service

import (
	"context"
	"lang-bot/internal/translation"
	"lang-bot/internal/word"
	"time"

	"github.com/google/uuid"
)

type WordService struct {
	wordRepository     word.Repository
	translationService translation.Service
}

func NewWordService(wordRepository word.Repository, translationService translation.Service) WordService {
	return WordService{
		wordRepository:     wordRepository,
		translationService: translationService,
	}

}

func (s WordService) AddWord(ctx context.Context, text string) (word.Word, error) {
	translation, err := s.translationService.GetTranslation(ctx, text)
	if err != nil {
		return word.Word{}, err
	}

	newWord := word.Word{
		ID:          uuid.New(),
		Text:        text,
		Translation: translation,
		CreatedAt:   time.Now(),
	}

	err = s.wordRepository.AddWord(ctx, newWord)
	if err != nil {
		return word.Word{}, err
	}

	return newWord, nil
}

func (s WordService) GetWords(ctx context.Context) ([]word.Word, error) {
	words, err := s.wordRepository.GetWords(ctx)
	if err != nil {
		return nil, err
	}

	return words, nil
}
