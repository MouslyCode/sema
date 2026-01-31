package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadPDF(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file required"})
		return
	}

	if file.Header.Get("Content-Type") != "application/pdf" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only PDF allowed"})
		return
	}

	_ = os.MkdirAll("uploads", os.ModePerm)
	path := fmt.Sprintf("uploads/%s", file.Filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "upload success"})
}
