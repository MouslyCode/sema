package handler

import (
	charts "github.com/MouslyCode/sema/chart"
	"github.com/gin-gonic/gin"
)

func TradingChart(c *gin.Context) {
	chart, err := charts.GenerateTradingChart("SIMA.csv")
	if err != nil {
		c.String(500, err.Error())
		return
	}
	chart.Render(c.Writer)
}
