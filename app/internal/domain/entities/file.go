package entities

import (
	"app/internal/domain/repositories"
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID           uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Generation   uint
	Transactions []*Transaction
}

func (s *File) OnCreate(uuidgen repositories.UUIDGenerator, clock repositories.Clock) error {
	id, err := uuidgen.NewV7()
	if err != nil {
		return err
	}
	s.ID = id
	s.CreatedAt = clock.Now().Round(time.Second)
	s.UpdatedAt = clock.Now().Round(time.Second)
	s.Generation = 1
	for _, transaction := range s.Transactions {
		err := transaction.OnCreate(uuidgen, clock)
		if err != nil {
			return err
		}
	}

	return nil
}
