package middleware

import (
	"github.com/gin-gonic/gin"
	"golang_server/models"
	"log"
	"net/http"
	"time"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.Use(logger)
	r.Use(errorHandler)

	v4 := r.Group("/v4")
	{
		v4.GET("/events", getEvents)
	}

	/*r.GET("/events", func(c *gin.Context) {

	})*/

	return r
}

func getEvents(c *gin.Context) {
	log.Printf("Current time: %s\n", c.MustGet("time").(time.Time))

	events, err := models.Events().All(Ctx, db.DbInstance)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, events)
}
