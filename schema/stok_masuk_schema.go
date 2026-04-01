package schema

type CreateStockInRequest struct {
	Supplier string        `json:"supplier"`
	Items     []StockInItem `json:"items"`
}

type StockInItem struct {
	ProductID int     `json:"product_id"`
	Qty       float64 `json:"qty"`
}

type UpdateStockInStatusRequest struct {
	Status string `json:"status"`
}