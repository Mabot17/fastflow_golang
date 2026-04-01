package model

type StockOut struct {
	ID        int
	Status    string
	Pelanggan string
	CreatedAt string
	UpdatedAt string
}

type StockOutProduct struct {
	ID         int
	StockOutID int
	ProductID  int
	Qty        float64
}