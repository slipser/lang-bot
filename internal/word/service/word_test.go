package service_test

import (
	"context"
	"lang-bot/internal/word/repository/localcache"
	"lang-bot/internal/word/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWord(t *testing.T) {
	ls := localcache.NewWordLocalStorage()
	s := service.NewWordService(ls)

	text := "Test"

	w, err := s.AddWord(context.Background(), text)
	assert.NoError(t, err)

	assert.Equal(t, text, w.Text)
}

func TestGetWords(t *testing.T) {
	ls := localcache.NewWordLocalStorage()
	s := service.NewWordService(ls)

	count := 10

	for i := 0; i < count; i++ {
		_, err := s.AddWord(context.Background(), "Test")
		assert.NoError(t, err)
	}

	w, err := s.GetWords(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, len(w), count)
}
