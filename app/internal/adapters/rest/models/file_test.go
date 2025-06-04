package models_test

import (
	"app/internal/adapters/rest/models"
	"app/internal/domain/entities"
	"context"
	"log/slog"
	"reflect"
	"testing"
)

func TestFileUpload_ToDomain(t *testing.T) {
	type fields struct {
		Transactions []models.TransactionCreate
	}
	type args struct {
		ctx    context.Context
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.File
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				Transactions: []models.TransactionCreate{
					{
						UserID:    testUUID.String(),
						Amount:    testAmount,
						Currency:  "USD",
						Timestamp: testTimestampString,
					},
				},
			},
			args: args{
				ctx:    context.TODO(),
				logger: slog.Default(),
			},
			want: &entities.File{
				Transactions: []*entities.Transaction{
					{
						UserID:    testUUID,
						Amount:    testAmount,
						Currency:  entities.USD,
						Timestamp: testTimestamp,
					},
				},
			},
		},
		{
			name: "unhappy path: empty transactions list",
			fields: fields{
				Transactions: []models.TransactionCreate{},
			},
			args: args{
				ctx:    context.TODO(),
				logger: slog.Default(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &models.FileUpload{
				Transactions: tt.fields.Transactions,
			}
			got, err := s.ToDomain(tt.args.ctx, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileUpload.ToDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileUpload.ToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
