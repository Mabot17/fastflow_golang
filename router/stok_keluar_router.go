package router

import (
	"github.com/gin-gonic/gin"

	"inventaris-app/crud"
	"inventaris-app/schema"
)
// =======================
// GET LIST
// =======================

// @Summary Get Stock Out List
// @Tags Stock Out
// @Produce json
// @Success 200 {array} model.StockOut
// @Router /stock-out [get]
func GetStockOutListHandler(c *gin.Context) {

	data, err := crud.GetStockOutList()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// =======================
// GET DETAIL
// =======================

// @Summary Get Stock Out Detail
// @Tags Stock Out
// @Produce json
// @Param id path string true "Stock Out ID"
// @Success 200 {array} model.StockOut
// @Router /stock-out/{id} [get]
func GetStockOutDetailHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := crud.GetStockOutByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// =======================
// CREATE (DRAFT / ALLOCATE)
// =======================

// @Summary Allocate Stock Out (Reserve)
// @Description Create stock out draft and reserve stock (does not reduce physical stock yet)
// @Tags Stock Out
// @Accept json
// @Produce json
// @Param data body schema.CreateStockOutRequest true "Stock Out Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock-out [post]
func CreateStockOutHandler(c *gin.Context) {
	var req schema.CreateStockOutRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := crud.CreateStockOut(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "allocated"})
}

// =======================
// CANCEL (ROLLBACK RESERVE)
// =======================

// @Summary Cancel Stock Out
// @Description Cancel stock out and rollback reserved stock
// @Tags Stock Out
// @Produce json
// @Param id path string true "Stock Out ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock-out/{id}/cancel [post]
func CancelStockOutHandler(c *gin.Context) {
	id := c.Param("id")

	err := crud.CancelStockOut(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "cancelled"})
}

// =======================
// DONE (FINAL COMMIT)
// =======================

// @Summary Finish Stock Out
// @Description Finalize stock out and reduce physical stock
// @Tags Stock Out
// @Produce json
// @Param id path string true "Stock Out ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /stock-out/{id}/finish [post]
func FinishStockOutHandler(c *gin.Context) {
	id := c.Param("id")

	err := crud.FinishStockOut(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "done"})
}

// =======================
// REGISTER
// =======================

func RegisterStockOutRoutes(r *gin.RouterGroup) {
	r.POST("/stock-out", CreateStockOutHandler)
	r.POST("/stock-out/:id/cancel", CancelStockOutHandler)
	r.POST("/stock-out/:id/finish", FinishStockOutHandler)
	r.GET("/stock-out", GetStockOutListHandler)
	r.GET("/stock-out/:id", GetStockOutDetailHandler)
}