package handlers

import (
	"app/internal/adapters/rest/models"
	"app/internal/domain/entities"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type UploadFile struct {
	backupLogger *slog.Logger
	fileUploader FileUploader
}

type FileUploader interface {
	Execute(ctx context.Context, file *entities.File) (string, error)
}

func (s *UploadFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file models.FileUpload

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		s.backupLogger.ErrorContext(
			r.Context(),
			"failed to parse file upload request",
			"error", err,
		)
		http.Error(w, "request body is invalid", http.StatusBadRequest)
		return
	}

	domainFile, err := file.ToDomain(r.Context(), s.backupLogger)
	if err != nil {
		s.backupLogger.ErrorContext(r.Context(),
			"failed to convert file upload request to domain",
			"error", err,
		)
		http.Error(w, "request body is invalid", http.StatusBadRequest)
		return
	}

	uri, err := s.fileUploader.Execute(r.Context(), domainFile)
	if err != nil {
		switch err.(type) {
		default:
			s.backupLogger.ErrorContext(r.Context(), "unknown error while uploading file", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uri)
}
