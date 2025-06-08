package handlers_test

import (
	"app/internal/adapters/rest/handlers"
	"app/internal/adapters/rest/models"
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile_ServeHTTP(t *testing.T) {
	var (
		testUUID                  = uuid.New()
		testAmount          int64 = 10
		testTimestampString       = "2006-01-02T15:04:05Z"
		// testTimestamp, _          = time.Parse(time.RFC3339, testTimestampString)
	)

	type fields struct {
		logger *slog.Logger
	}
	type args struct {
		w *httptest.ResponseRecorder
		b any
	}
	type want struct {
		statusCode int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   want
	}{
		{
			name: "happy path",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				w: httptest.NewRecorder(),
				b: &models.FileUpload{
					Transactions: []models.TransactionCreate{
						{
							UserID:    testUUID.String(),
							Amount:    testAmount,
							Currency:  "USD",
							Timestamp: testTimestampString,
						},
					},
				},
			},
			want: want{
				statusCode: http.StatusCreated,
			},
		},
		{
			name: "unhappy path: invalid request body",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				w: httptest.NewRecorder(),
				b: "boom",
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
		{
			name: "unhappy path: invalid file",
			fields: fields{
				logger: slog.Default(),
			},
			args: args{
				w: httptest.NewRecorder(),
				b: &models.FileUpload{},
			},
			want: want{
				statusCode: http.StatusBadRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			requestBody, _ := json.Marshal(tt.args.b)
			testReq, _ := http.NewRequest(http.MethodPost, "/file", bytes.NewBuffer(requestBody))
			handler := handlers.NewUploadFile(tt.fields.logger)

			handler.ServeHTTP(tt.args.w, testReq)

			assert.Equal(t, tt.want.statusCode, tt.args.w.Code)
		})
	}
}
