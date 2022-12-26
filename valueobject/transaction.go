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

func NewTransaction(amount int, fom uuid.UUID, to uuid.UUID) Transaction {
	return Transaction{
		amount:    amount,
		fom:       fom,
		to:        to,
		createdAt: time.Now(),
	}
}
