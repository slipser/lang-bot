package word

import "context"

type Repository interface {
	AddWord(ctx context.Context, word Word) error
	GetWords(ctx context.Context) ([]Word, error)
}
