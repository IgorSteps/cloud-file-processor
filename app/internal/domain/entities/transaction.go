package entities

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdateAt   time.Time
	Generation uint

	UserID    uuid.UUID
	Amount    float64
	Currency  Currency
	Timestamp time.Time
}
