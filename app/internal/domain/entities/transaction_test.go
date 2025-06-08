package entities_test

import (
	"app/internal/domain/entities"
	"app/internal/domain/repositories"
	mocks_repositories "app/mocks/repositories"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	testUUID, _         = uuid.NewV7()
	testTimestampString = "2006-01-02T15:04:05Z"
	testTimestamp, _    = time.Parse(time.RFC3339, testTimestampString)
	errorTest           = errors.New("boom")
)

func TestTransaction_OnCreate(t *testing.T) {
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
			args: args{
				uuidgen: newUnhappyUUIDGenerator(t),
				clock:   newNoopClock(t),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expectedTransaction := &entities.Transaction{
				ID:         tt.expected.ID,
				CreatedAt:  tt.expected.CreatedAt,
				UpdatedAt:  tt.expected.UpdatedAt,
				Generation: tt.expected.Generation,
			}
			actualTransaction := &entities.Transaction{}

			err := actualTransaction.OnCreate(tt.args.uuidgen, tt.args.clock)

			if !tt.wantErr {
				assert.NoError(t, err)
				assert.Equal(t, expectedTransaction, actualTransaction)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func newHappyUUIDGenerator(t *testing.T) repositories.UUIDGenerator {
	uuidGen := mocks_repositories.NewUUIDGenerator(t)
	uuidGen.EXPECT().NewV7().Return(testUUID, nil).Once()
	return uuidGen
}

func newUnhappyUUIDGenerator(t *testing.T) repositories.UUIDGenerator {
	uuidGen := mocks_repositories.NewUUIDGenerator(t)
	uuidGen.EXPECT().NewV7().Return(uuid.Nil, errorTest).Once()
	return uuidGen
}

func newHappyClock(t *testing.T) repositories.Clock {
	clock := mocks_repositories.NewClock(t)
	clock.EXPECT().Now().Return(testTimestamp).Once()
	clock.EXPECT().Now().Return(testTimestamp).Once()
	return clock
}

func newNoopClock(t *testing.T) repositories.Clock {
	clock := mocks_repositories.NewClock(t)
	return clock
}
