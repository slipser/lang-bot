package word

import "context"

type Service interface {
	AddWord(ctx context.Context, text string) (Word, error)
}
