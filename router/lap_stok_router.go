package router

import (
	"github.com/gin-gonic/gin"
	"inventaris-app/crud"
	"inventaris-app/schema"
)

// @Summary Get Stock Report
// @Description Get stock movement report
// @Tags Report
// @Produce json
// @Param product_name query string false "Product Name"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {array} model.LapStock
// @Router /lap-stok [get]
func GetLapStockHandler(c *gin.Context) {
	var req schema.GetLapStockRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := crud.GetLapStock(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, data)
}

// @Summary Export Stock Report Excel
// @Description Export laporan stok ke Excel
// @Tags Report
// @Produce application/octet-stream
// @Param product_name query string false "Product Name"
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Success 200 {file} file
// @Router /lap-stok/export [get]
func ExportLapStockHandler(c *gin.Context) {

	var req schema.GetLapStockRequest

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	buf, err := crud.ExportLapStockExcel(req)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// 🔥 kirim sebagai file download
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=lap_stok.xlsx")

	c.Data(200, "application/octet-stream", buf.Bytes())
}

func RegisterLapStockRoutes(r *gin.RouterGroup) {
	r.GET("/lap-stok", GetLapStockHandler)
	r.GET("/lap-stok/export", ExportLapStockHandler)
}