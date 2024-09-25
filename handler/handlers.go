package handler

import (
	"github.com/Latinaxia/url-shorter/shortener"
	"github.com/Latinaxia/url-shorter/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortUrl(c *gin.Context) {
	/// Implementation to be added
	var req UrlCreationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	shortUrl := shortener.GenerateShortLink(req.LongUrl, req.UserId)
	store.SaveUrlMapping(shortUrl, req.LongUrl, req.UserId)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})

}

func HandleShortUrlRedirect(c *gin.Context) {
	/// Implementation to be added
	shortUrl := c.Param("shortUrl")
	finalurl := store.RetrieveInitialUrl(shortUrl)
	c.Redirect(302, finalurl)
}
