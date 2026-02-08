package models

type DailySalesReport struct {
	TotalRevenue   int                `json:"total_revenue"`
	TotalTransaksi int                `json:"total_transaksi"`
	ProdukTerlaris BestSellingProduct `json:"produk_terlaris"`
}

type BestSellingProduct struct {
	Nama       string `json:"nama"`
	QtyTerjual int    `json:"qty_terjual"`
}

type RangeReportItem struct {
	Tanggal        string             `json:"tanggal"`
	TotalRevenue   int                `json:"total_revenue"`
	TotalTransaksi int                `json:"total_transaksi"`
	ProdukTerlaris BestSellingProduct `json:"produk_terlaris"`
}
