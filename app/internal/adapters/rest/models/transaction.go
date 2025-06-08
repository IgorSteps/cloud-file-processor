package models

import (
	"app/internal/domain/entities"
	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type TransactionCreate struct {
	UserID    string `json:"userId"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Timestamp string `json:"timestamp"`
}

func (t *TransactionCreate) ToDomain(ctx context.Context, logger *slog.Logger) (*entities.Transaction, error) {
	userID, err := uuid.Parse(t.UserID)
	if err != nil {
		logger.ErrorContext(ctx,
			"failed to parse user id",
			"user_id", t.UserID,
			"error", err,
		)

		return nil, entities.NewInvalidInputError("failed to parse user id")
	}

	currency, err := entities.ParseCurrency(t.Currency)
	if err != nil {
		logger.ErrorContext(
			ctx,
			"failed to parse currency",
			"currency", t.Currency,
			"error", err)

		return nil, entities.NewInvalidInputError("failed to parse currency")
	}

	timestamp, err := time.Parse(time.RFC3339, t.Timestamp)
	if err != nil {
		logger.ErrorContext(ctx,
			"failed to parse timestamp",
			"timestamp", t.Timestamp,
			"error", err,
		)

		return nil, entities.NewInvalidInputError("failed to parse timestamp")
	}

	return &entities.Transaction{
		UserID:    userID,
		Amount:    t.Amount,
		Currency:  currency,
		Timestamp: timestamp,
	}, nil
}
