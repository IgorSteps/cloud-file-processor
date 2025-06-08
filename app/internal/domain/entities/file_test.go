package entities_test

import (
	"app/internal/domain/entities"
	"app/internal/domain/repositories"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestFile_OnCreate(t *testing.T) {
	type expected struct {
		ID         uuid.UUID
		CreatedAt  time.Time
		UpdatedAt  time.Time
		Generation uint
	}
	type args struct {
		uuidgen repositories.UUIDGenerator
		clock   repositories.Clock
	}
	tests := []struct {
		name     string
		expected expected
		args     args
		wantErr  bool
	}{
		{
			name: "happy path",
			expected: expected{
				ID:         testUUID,
				CreatedAt:  testTimestamp,
				UpdatedAt:  testTimestamp,
				Generation: 1,
			},
			args: args{
				uuidgen: newHappyUUIDGenerator(t),
				clock:   newHappyClock(t),
			},
			wantErr: false,
		},
		{
			name: "unhappy path",
			expected: expected{
				ID:         testUUID,
				CreatedAt:  testTimestamp,
				UpdatedAt:  testTimestamp,
				Generation: 1,
			},
			args: args{
				uuidgen: newUnhappyUUIDGenerator(t),
				clock:   newNoopClock(t),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedFile := &entities.File{
				ID:         tt.expected.ID,
				CreatedAt:  tt.expected.CreatedAt,
				UpdatedAt:  tt.expected.UpdatedAt,
				Generation: tt.expected.Generation,
			}

			actualFile := &entities.File{}

			err := actualFile.OnCreate(tt.args.uuidgen, tt.args.clock)

			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, expectedFile, actualFile)
			} else {
				assert.Error(t, err)
			}

		})
	}
}
