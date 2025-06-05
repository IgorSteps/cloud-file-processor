package usecases

import (
	"app/internal/domain/entities"
	"log/slog"
)

type UploadFile struct {
	logger *slog.Logger
}

func (s *UploadFile) Execute(file *entities.File) (string, error) {
	// TODO

	return "hi", nil
}
