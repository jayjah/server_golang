package routes

import (
	"github.com/gin-gonic/gin"
	"golang_server/database"
	"net/http"
	"time"
)

func Ping(c *gin.Context) {
	var dbTime time.Time
	database.Db.Instance.Raw("SELECT now();").Find(&dbTime)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"time":    dbTime,
	})
}
