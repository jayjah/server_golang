package routes

import (
	"github.com/gin-gonic/gin"
	"golang_server/database"
	"golang_server/gorm_models/model"
	"log"
	"net/http"
	"time"
)

func GetEvents(c *gin.Context) {
	log.Printf("Current time: %s\n", c.MustGet("time").(time.Time))

	var events []model.Event
	database.Db.Instance.Model(&model.Event{}).Preload("Image").Find(&events)
	c.JSON(http.StatusOK, events)
}
