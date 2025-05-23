package handlers

import (
	"app/internal/adapters/rest/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

type UploadFile struct {
	backupLogger *slog.Logger
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

}
