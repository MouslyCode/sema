package router

import (
	"github.com/MouslyCode/sema/handler"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/news", handler.GetNews)
	r.GET("/charts/trading", handler.TradingChart)
	r.POST("/upload/pdf", handler.UploadPDF)
}
