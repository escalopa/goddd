package valueobject

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	amount    int
	fom       uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
