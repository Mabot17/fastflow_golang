package model
import "time"

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

type StockOutList struct {
	ID          int       `json:"id"`
	Pelanggan   string    `json:"pelanggan"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`

	ProductID   int     `json:"product_id"`
	ProductName string  `json:"product_name"`
	Qty         float64 `json:"qty"`
}