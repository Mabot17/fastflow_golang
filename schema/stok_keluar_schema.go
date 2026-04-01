package schema

type CreateStockOutRequest struct {
	Pelanggan string         `json:"pelanggan"`
	Items     []StockOutItem `json:"items"`
}

type StockOutItem struct {
	ProductID int     `json:"product_id"`
	Qty       float64 `json:"qty"`
}