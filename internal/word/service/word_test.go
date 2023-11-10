package service_test

import (
	"context"
	translationService "lang-bot/internal/translation/service"
	"lang-bot/internal/word/repository/localcache"
	"lang-bot/internal/word/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWord(t *testing.T) {
	ls := localcache.NewWordLocalStorage()
	ts := translationService.NewTranslationService()
	s := service.NewWordService(ls, ts)

	text := "Test"

	w, err := s.AddWord(context.Background(), text)
	assert.NoError(t, err)

	assert.Equal(t, text, w.Text)
}

func TestGetWords(t *testing.T) {
	ls := localcache.NewWordLocalStorage()
	ts := translationService.NewTranslationService()
	s := service.NewWordService(ls, ts)

	count := 10

	for i := 0; i < count; i++ {
		_, err := s.AddWord(context.Background(), "Test")
		assert.NoError(t, err)
	}

	w, err := s.GetWords(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, len(w), count)
}
