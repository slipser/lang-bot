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
