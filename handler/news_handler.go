package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type NewsAPIResponse struct {
	Status   string `json:"status"`
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		PublishedAt string `json:"publishedAt"`
		Source      struct {
			Name string `json:"name"`
		} `json:"source"`
	} `json:"articles"`
}

func GetNews(c *gin.Context) {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "API_KEY not set",
		})
		return
	}

	url := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=bitcoin&apiKey=%s",
		apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var result NewsAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, result.Articles)
}
