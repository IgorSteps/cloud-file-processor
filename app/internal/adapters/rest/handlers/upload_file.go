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
	logger *slog.Logger
	//fileUploader FileUploader
}

func NewUploadFile(logger *slog.Logger) *UploadFile {
	return &UploadFile{
		logger: logger,
	}
}

type FileUploader interface {
	Execute(ctx context.Context, file *entities.File) (string, error)
}

func (s *UploadFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var file models.FileUpload

	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		s.logger.ErrorContext(
			r.Context(),
			"failed to parse file upload request",
			"error", err,
		)
		http.Error(w, "request body is invalid", http.StatusBadRequest)
		return
	}

	_, err = file.ToDomain(r.Context(), s.logger)
	if err != nil {
		s.logger.ErrorContext(r.Context(),
			"failed to convert file upload request to domain",
			"error", err,
		)
		http.Error(w, "file is invalid", http.StatusBadRequest)
		return
	}

	// uri, err := s.fileUploader.Execute(r.Context(), domainFile)
	// if err != nil {
	// 	switch err.(type) {
	// 	default:
	// 		s.logger.ErrorContext(r.Context(), "unknown error while uploading file", "error", err)
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// }

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode("hi")
	if err != nil {
		s.logger.ErrorContext(r.Context(),
			"failed to encode response",
			"response", "TODO",
			"error", err,
		)
		http.Error(w, "response body is invalid", http.StatusInternalServerError)
		return
	}
}
