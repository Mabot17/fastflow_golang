package schema

type CreateProductRequest struct {
	Name string `json:"name" binding:"required"`
	SKU  string `json:"sku" binding:"required"`
}

type UpdateProductRequest struct {
	Name string `json:"name"`
	SKU  string `json:"sku"`
}

type ProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	SKU  string `json:"sku"`
}