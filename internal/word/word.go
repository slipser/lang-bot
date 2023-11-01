package word

import (
	"time"

	"github.com/google/uuid"
)

type Word struct {
	ID   uuid.UUID
	Text string
	//Translation string
	CreatedAt time.Time
}
