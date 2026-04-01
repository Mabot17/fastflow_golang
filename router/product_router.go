package router

import (
	"github.com/gin-gonic/gin"

	"inventaris-app/crud"
	"inventaris-app/schema"
	"inventaris-app/model"
)

// hack biar kepake
var _ = model.Product{}

// =======================
// HANDLERS (dibaca Swag)
// =======================

// @Summary Get all products
// @Description Get all products
// @Tags Products
// @Produce json
// @Success 200 {array} model.Product
// @Router /products [get]
func GetProducts(c *gin.Context) {
	data, err := crud.GetAllProducts()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, data)
}

// @Summary Get product by ID
// @Description Get single product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} model.Product
// @Failure 404 {object} map[string]string
// @Router /products/{id} [get]
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	data, err := crud.GetProductByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Not found"})
		return
	}
	c.JSON(200, data)
}

// @Summary Create product
// @Description Create new product
// @Tags Products
// @Accept json
// @Produce json
// @Param data body schema.CreateProductRequest true "Create Product"
// @Success 200 {object} model.Product
// @Failure 400 {object} map[string]string
// @Router /products [post]
func CreateProduct(c *gin.Context) {
	var req schema.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := crud.CreateProduct(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// @Summary Update product (full)
// @Description Full update product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param data body schema.UpdateProductRequest true "Update Product"
// @Success 200 {object} model.Product
// @Router /products/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req schema.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := crud.UpdateProduct(id, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// @Summary Patch product (partial)
// @Description Partial update product
// @Tags Products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param data body schema.UpdateProductRequest true "Patch Product"
// @Success 200 {object} model.Product
// @Router /products/{id} [patch]
func PatchProduct(c *gin.Context) {
	id := c.Param("id")
	var req schema.UpdateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := crud.PatchProduct(id, req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// @Summary Delete product
// @Description Delete product by ID
// @Tags Products
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} map[string]string
// @Router /products/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	err := crud.DeleteProduct(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "deleted"})
}

// =======================
// ROUTER
// =======================

func RegisterProductRoutes(r *gin.RouterGroup) {

	r.GET("/products", GetProducts)
	r.GET("/products/:id", GetProductByID)

	r.POST("/products", CreateProduct)

	r.PUT("/products/:id", UpdateProduct)
	r.PATCH("/products/:id", PatchProduct)

	r.DELETE("/products/:id", DeleteProduct)
}