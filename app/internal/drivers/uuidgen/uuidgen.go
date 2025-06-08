package uuidgen

import "github.com/google/uuid"

type UUIDGenerator struct{}

func (s *UUIDGenerator) NewV7() (uuid.UUID, error) {
	return uuid.NewV7()
}
