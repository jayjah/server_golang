package routes

import (
	"github.com/gin-gonic/gin"
	"golang_server/middleware"
	"golang_server/models"
	"log"
	"net/http"
	"time"
)

func GetEvents(c *gin.Context) {
	log.Printf("Current time: %s\n", c.MustGet("time").(time.Time))

	events, err := models.Events().All(middleware.Ctx, nil)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, events)
}
