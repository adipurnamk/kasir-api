package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
	"strings"
)

// TransactionRepository handles transaction-related database operations.
type TransactionRepository struct {
	db *sql.DB
}

// NewTransactionRepository creates a new TransactionRepository instance
func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Batch insert details in a single multi-row INSERT to avoid N round-trips
	if len(details) > 0 {
		valueStrings := make([]string, 0, len(details))
		valueArgs := make([]interface{}, 0, len(details)*4)

		for i := range details {
			details[i].TransactionID = transactionID
			// placeholders for postgres: $1, $2, ...
			idx := i * 4
			valueStrings = append(valueStrings, fmt.Sprintf("($%d,$%d,$%d,$%d)", idx+1, idx+2, idx+3, idx+4))
			valueArgs = append(valueArgs, transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		}

		stmt := fmt.Sprintf("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES %s", strings.Join(valueStrings, ","))
		_, err = tx.Exec(stmt, valueArgs...)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
