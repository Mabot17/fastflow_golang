package schema

type GetLapStockRequest struct {
	ProductName string `form:"product_name"`
	Sumber      string `form:"sumber"` // customer / supplier
	StartDate   string `form:"start_date"`
	EndDate     string `form:"end_date"`
}