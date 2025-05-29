package models

import (
	"app/internal/domain/entities"
	"context"
	"log/slog"
)

type FileUpload struct {
	Transactions []TransactionCreate `json:"transactions"`
}

func (s *FileUpload) ToDomain(ctx context.Context, logger *slog.Logger) (*entities.File, error) {
	var domainTransactions []*entities.Transaction

	if len(s.Transactions) == 0 {
		return nil, entities.NewInvalidInputError("empty transactions list")
	}

	for _, transaction := range s.Transactions {
		domainTran, err := transaction.ToDomain(ctx, logger)
		if err != nil {
			return nil, err
		}
		domainTransactions = append(domainTransactions, domainTran)
	}

	return &entities.File{
		Transactions: domainTransactions,
	}, nil
}
