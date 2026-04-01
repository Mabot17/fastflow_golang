package router

import (
	"github.com/gin-gonic/gin"

	"inventaris-app/crud"
	"inventaris-app/schema"
)

// =======================
// GET LIST
// =======================

// @Summary Get Stock In List
// @Tags Stock In
// @Produce json
// @Success 200 {array} model.StockIn
// @Router /stock-in [get]
func GetStockInListHandler(c *gin.Context) {

	data, err := crud.GetStockInList()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// =======================
// GET DETAIL
// =======================

// @Summary Get Stock In Detail
// @Tags Stock In
// @Produce json
// @Param id path string true "Stock In ID"
// @Success 200 {object} model.StockIn
// @Router /stock-in/{id} [get]
func GetStockInDetailHandler(c *gin.Context) {
	id := c.Param("id")

	data, err := crud.GetStockInByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// =======================
// CREATE
// =======================

// @Summary Create Stock In (Multi Product)
// @Tags Stock In
// @Accept json
// @Produce json
// @Param data body schema.CreateStockInRequest true "Stock In"
// @Success 200 {object} map[string]string
// @Router /stock-in [post]
func CreateStockInHandler(c *gin.Context) {
	var req schema.CreateStockInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := crud.CreateStockIn(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "created"})
}

// =======================
// FINISH (DONE)
// =======================

// @Summary Finish Stock In
// @Tags Stock In
// @Produce json
// @Param id path string true "Stock In ID"
// @Success 200 {object} map[string]string
// @Router /stock-in/{id}/finish [post]
func FinishStockInHandler(c *gin.Context) {
	id := c.Param("id")

	err := crud.FinishStockIn(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "done"})
}

// =======================
// CANCEL
// =======================

// @Summary Cancel Stock In
// @Tags Stock In
// @Produce json
// @Param id path string true "Stock In ID"
// @Success 200 {object} map[string]string
// @Router /stock-in/{id}/cancel [post]
func CancelStockInHandler(c *gin.Context) {
	id := c.Param("id")

	err := crud.CancelStockIn(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "cancelled"})
}

// =======================
// REGISTER
// =======================

func RegisterStockInRoutes(r *gin.RouterGroup) {
	r.GET("/stock-in", GetStockInListHandler)
	r.GET("/stock-in/:id", GetStockInDetailHandler)
	r.POST("/stock-in", CreateStockInHandler)
	r.POST("/stock-in/:id/finish", FinishStockInHandler)
	r.POST("/stock-in/:id/cancel", CancelStockInHandler)
}