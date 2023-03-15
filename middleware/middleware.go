package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func logger(c *gin.Context) {
	// before request

	t := time.Now()
	c.Set("time", t)

	// handle request
	c.Next()

	// after request

	latency := time.Since(t)
	log.Printf("Latency: %s", latency)

	status := c.Writer.Status()
	log.Printf("Request StatusCode: %d", status)
}

func errorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors.String()) <= 0 && len(c.Errors) < 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	for _, err := range c.Errors {
		log.Print(err)
	}

	if len(c.Errors) > 0 {
		c.JSON(-1, c.Errors.JSON())
	}
}
