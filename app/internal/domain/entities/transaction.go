package entities

import (
	"github.com/google/uuid"
)

type Transaction struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Amount    float64
	Currency  Currency
	Timestamp Clock
}
