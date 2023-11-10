package translation

import "context"

type Service interface {
	GetTranslation(ctx context.Context, text string) (string, error)
}
