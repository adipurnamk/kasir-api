package repositories

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetByID_Returns_Product_With_Category(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %v", err)
	}
	defer db.Close()

	repo := NewProductRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "price", "stock", "category_id", "category_name", "category_description"}).
		AddRow(1, "Prod A", 10000, 5, 2, "Minuman", "Menghilangkan dahaga")

	mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT p.id, p.name, p.price, p.stock, p.category_id, c.name AS category_name, c.description AS category_description
                FROM products p
                LEFT JOIN categories c ON p.category_id = c.id
                WHERE p.id = $1`)).
		WithArgs(1).
		WillReturnRows(rows)

	prod, err := repo.GetByID(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if prod.Category == nil {
		t.Fatalf("expected category, got nil")
	}
	if prod.Category.Name != "Minuman" {
		t.Fatalf("expected category name Minuman, got %s", prod.Category.Name)
	}
	if prod.Category.Description != "Menghilangkan dahaga" {
		t.Fatalf("expected category description Menghilangkan dahaga, got %s", prod.Category.Description)
	}
}
