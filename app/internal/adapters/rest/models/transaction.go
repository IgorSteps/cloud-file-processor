package models

import (
	"app/internal/domain/entities"
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type TransactionCreate struct {
	UserID    string  `json:"userId"`
	Amount    float64 `json:"amount"`
	Currency  string  `json:"currency"`
	Timestamp string  `json:"timestamp"`
}

func (t *TransactionCreate) ToDomain(ctx context.Context, logger *slog.Logger) (*entities.Transaction, error) {
	userID, err := uuid.Parse(t.UserID)
	if err != nil {
		logger.ErrorContext(ctx,
			"invalid user id",
			"user_id", t.UserID,
			"error", err,
		)

		return nil, err
	}

	if t.Amount < 0 {
		logger.ErrorContext(ctx,
			"invalid amount",
			"amount", t.Amount,
			"error", "cannot be negative",
		)

		return nil, errors.New("amount cannot be negative")
	}

	currency, err := entities.ParseCurrency(t.Currency)
	if err != nil {
		logger.ErrorContext(
			ctx,
			"invalid currency",
			"currency", t.Currency,
			"error", err)

		return nil, err
	}

	timestamp, err := time.Parse(time.RFC3339, t.Currency)
	if err != nil {
		logger.ErrorContext(ctx,
			"invalid timestamp",
			"timestamp", t.Timestamp,
			"error", err,
		)

		return nil, err
	}

	return &entities.Transaction{
		UserID:    userID,
		Amount:    t.Amount,
		Currency:  currency,
		Timestamp: timestamp,
	}, nil
}
