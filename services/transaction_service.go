package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

// TransactionService provides transaction-related operations.
type TransactionService struct {
	repo *repositories.TransactionRepository
}

// NewTransactionService creates a new TransactionService instance
func NewTransactionService(repo *repositories.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

func (s *TransactionService) Checkout(items []models.CheckoutItem, useLock bool) (*models.Transaction, error) {
	return s.repo.CreateTransaction(items)
}
