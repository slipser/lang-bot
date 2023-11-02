package localcache_test

import (
	"context"
	"lang-bot/internal/word"
	"lang-bot/internal/word/repository/localcache"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAddWord(t *testing.T) {
	ls := localcache.NewWordLocalStorage()

	w := word.Word{
		ID:        uuid.New(),
		Text:      "Test",
		CreatedAt: time.Now(),
	}

	err := ls.AddWord(context.Background(), w)
	assert.NoError(t, err)
}

func TestGetWords(t *testing.T) {
	ls := localcache.NewWordLocalStorage()

	count := 10

	for i := 0; i < count; i++ {
		w := word.Word{
			ID:        uuid.New(),
			Text:      "Test",
			CreatedAt: time.Now(),
		}
		err := ls.AddWord(context.Background(), w)
		assert.NoError(t, err)
	}

	words, err := ls.GetWords(context.Background())
	assert.NoError(t, err)

	assert.Equal(t, len(words), count)
}
