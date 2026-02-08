package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

// ReportService provides report-related operations.
type ReportService struct {
	repo *repositories.ReportRepository
}

// NewReportService creates a new ReportService instance
func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{repo: repo}
}

// GetDailySalesReport returns today's sales summary
func (s *ReportService) GetDailySalesReport() (*models.DailySalesReport, error) {
	return s.repo.GetDailySalesReport()
}

// GetSalesReportByDateRange returns sales summary for a date range
func (s *ReportService) GetSalesReportByDateRange(startDate, endDate string) ([]models.RangeReportItem, error) {
	return s.repo.GetSalesReportByDateRange(startDate, endDate)
}
