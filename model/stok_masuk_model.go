package model

type StockIn struct {
	ID        int
	Status    string
	CreatedAt string
	UpdatedAt string
}

type StockInProduct struct {
	ID        int
	StockInID int
	ProductID int
	Qty       float64
}