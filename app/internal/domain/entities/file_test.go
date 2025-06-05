package entities_test

import (
	"app/internal/domain/entities"
	"app/internal/domain/repositories"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestFile_OnCreate(t *testing.T) {
	type fields struct {
		ID           uuid.UUID
		CreatedAt    time.Time
		UpdatedAt    time.Time
		Generation   uint
		Transactions []*entities.Transaction
	}
	type args struct {
		uuidgen repositories.UUIDGenerator
		clock   repositories.Clock
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &entities.File{
				ID:           tt.fields.ID,
				CreatedAt:    tt.fields.CreatedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				Generation:   tt.fields.Generation,
				Transactions: tt.fields.Transactions,
			}
			if err := s.OnCreate(tt.args.uuidgen, tt.args.clock); (err != nil) != tt.wantErr {
				t.Errorf("File.OnCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
