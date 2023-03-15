package routes

import (
	"github.com/gin-gonic/gin"
	"golang_server/database"
	"golang_server/models"
	"log"
	"net/http"
	"time"
)

func GetEvents(c *gin.Context) {
	log.Printf("Current time: %s\n", c.MustGet("time").(time.Time))

	events, err := models.Events().All(database.Ctx, database.Db.DbInstance)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, events)
}
