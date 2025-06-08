package models_test

import (
	"app/internal/adapters/rest/models"
	"app/internal/domain/entities"
	"context"
	"log/slog"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
)

var (
	testUUID                  = uuid.New()
	testAmount          int64 = 10
	testTimestampString       = "2006-01-02T15:04:05Z"
	testTimestamp, _          = time.Parse(time.RFC3339, testTimestampString)
)

func TestTransactionCreate_ToDomain(t *testing.T) {
	type fields struct {
		UserID    string
		Amount    int64
		Currency  string
		Timestamp string
	}
	type args struct {
		ctx    context.Context
		logger *slog.Logger
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Transaction
		wantErr bool
	}{
		{
			name: "happy path",
			fields: fields{
				UserID:    testUUID.String(),
				Amount:    10,
				Currency:  "USD",
				Timestamp: testTimestampString,
			},
			args: args{
				ctx:    context.TODO(),
				logger: slog.Default(),
			},
			want: &entities.Transaction{
				UserID:    testUUID,
				Amount:    testAmount,
				Currency:  entities.USD,
				Timestamp: testTimestamp,
			},
			wantErr: false,
		},
		{
			name: "invalid user id",
			fields: fields{
				UserID:    "boom",
				Amount:    10,
				Currency:  "USD",
				Timestamp: testTimestampString,
			},
			args: args{
				ctx:    context.TODO(),
				logger: slog.Default(),
			},
			wantErr: true,
		},
		{
			name: "invalid currency",
			fields: fields{
				UserID:    testUUID.String(),
				Amount:    10,
				Currency:  "boom",
				Timestamp: testTimestampString,
			},
			args: args{
				ctx:    context.TODO(),
				logger: slog.Default(),
			},
			wantErr: true,
		},
		{
			name: "invalid timestamp",
			fields: fields{
				UserID:    testUUID.String(),
				Amount:    10,
				Currency:  "USD",
				Timestamp: "boom",
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
			tr := &models.TransactionCreate{
				UserID:    tt.fields.UserID,
				Amount:    tt.fields.Amount,
				Currency:  tt.fields.Currency,
				Timestamp: tt.fields.Timestamp,
			}
			got, err := tr.ToDomain(tt.args.ctx, tt.args.logger)
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionCreate.ToDomain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TransactionCreate.ToDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}
