package router

import (
	"github.com/gin-gonic/gin"

	"inventaris-app/crud"
	"inventaris-app/schema"
)

func RegisterProductRoutes(r *gin.RouterGroup) {

	// GET ALL
	r.GET("/products", func(c *gin.Context) {
		data, err := crud.GetAllProducts()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, data)
	})

	// GET SINGLE
	r.GET("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		data, err := crud.GetProductByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "Not found"})
			return
		}
		c.JSON(200, data)
	})

	// CREATE
	r.POST("/products", func(c *gin.Context) {
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
	})

	// PUT (FULL UPDATE)
	r.PUT("/products/:id", func(c *gin.Context) {
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
	})

	// PATCH (PARTIAL)
	r.PATCH("/products/:id", func(c *gin.Context) {
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
	})

	// DELETE
	r.DELETE("/products/:id", func(c *gin.Context) {
		id := c.Param("id")

		err := crud.DeleteProduct(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "deleted"})
	})
}