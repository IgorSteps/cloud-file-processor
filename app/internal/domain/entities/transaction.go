package entities

import (
	"app/internal/domain/repositories"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Generation uint

	UserID    uuid.UUID
	Amount    int64
	Currency  Currency
	Timestamp time.Time
}

func (s *Transaction) OnCreate(uuidgen repositories.UUIDGenerator, clock repositories.Clock) error {
	id, err := uuidgen.NewV7()
	if err != nil {
		return err
	}
	s.ID = id
	s.CreatedAt = clock.Now().Round(time.Second)
	s.UpdatedAt = clock.Now().Round(time.Second)
	s.Generation = 1

	return nil
}
