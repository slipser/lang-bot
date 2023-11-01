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
