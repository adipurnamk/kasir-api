package repositories

import (
	"database/sql"
	"kasir-api/models"
	"time"
)

// ReportRepository handles report-related database operations.
type ReportRepository struct {
	db *sql.DB
}

// NewReportRepository creates a new ReportRepository instance
func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

// GetDailySalesReport returns today's sales summary
func (repo *ReportRepository) GetDailySalesReport() (*models.DailySalesReport, error) {
	report := &models.DailySalesReport{}

	// Get total revenue and transaction count for today
	today := time.Now().Format("2006-01-02")
	err := repo.db.QueryRow(`
		SELECT 
			COALESCE(SUM(total_amount), 0) as total_revenue,
			COUNT(*) as total_transaksi
		FROM transactions
		WHERE DATE(created_at) = $1
	`, today).Scan(&report.TotalRevenue, &report.TotalTransaksi)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	// Get best-selling product for today
	err = repo.db.QueryRow(`
		SELECT 
			p.name,
			SUM(td.quantity) as qty_terjual
		FROM transaction_details td
		JOIN products p ON td.product_id = p.id
		JOIN transactions t ON td.transaction_id = t.id
		WHERE DATE(t.created_at) = $1
		GROUP BY p.id, p.name
		ORDER BY qty_terjual DESC
		LIMIT 1
	`, today).Scan(&report.ProdukTerlaris.Nama, &report.ProdukTerlaris.QtyTerjual)

	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return report, nil
}

// GetSalesReportByDateRange returns sales summary for a date range
func (repo *ReportRepository) GetSalesReportByDateRange(startDate, endDate string) ([]models.RangeReportItem, error) {
	rows, err := repo.db.Query(`
		SELECT 
			DATE(t.created_at) as tanggal,
			COALESCE(SUM(t.total_amount), 0) as total_revenue,
			COUNT(DISTINCT t.id) as total_transaksi
		FROM transactions t
		WHERE DATE(t.created_at) >= $1 AND DATE(t.created_at) <= $2
		GROUP BY DATE(t.created_at)
		ORDER BY DATE(t.created_at) DESC
	`, startDate, endDate)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reports []models.RangeReportItem

	for rows.Next() {
		var item models.RangeReportItem
		err := rows.Scan(&item.Tanggal, &item.TotalRevenue, &item.TotalTransaksi)
		if err != nil {
			return nil, err
		}

		// Get best-selling product for this date
		err = repo.db.QueryRow(`
			SELECT 
				p.name,
				SUM(td.quantity) as qty_terjual
			FROM transaction_details td
			JOIN products p ON td.product_id = p.id
			JOIN transactions t ON td.transaction_id = t.id
			WHERE DATE(t.created_at) = $1
			GROUP BY p.id, p.name
			ORDER BY qty_terjual DESC
			LIMIT 1
		`, item.Tanggal).Scan(&item.ProdukTerlaris.Nama, &item.ProdukTerlaris.QtyTerjual)

		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		reports = append(reports, item)
	}

	return reports, rows.Err()
}
