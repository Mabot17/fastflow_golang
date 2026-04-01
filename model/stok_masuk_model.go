package model
import "time"

type StockIn struct {
	ID        int
	Status    string
	Supplier string
	CreatedAt string
	UpdatedAt string
}

type StockInProduct struct {
	ID        int
	StockInID int
	ProductID int
	Qty       float64
}

type StockInList struct {
	ID        int     `json:"id"`
	Supplier  string  `json:"supplier"`
	Status    string  `json:"status"`
	CreatedAt     time.Time

	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Qty         float64 `json:"qty"`
}