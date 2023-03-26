package routes

import (
	"github.com/gin-gonic/gin"
	"golang_server/database"
	"golang_server/gorm_models/model"
	"net/http"
)

func GetImages(c *gin.Context) {
	var images []model.Image
	database.Db.Instance.Find(&images)
	c.JSON(http.StatusOK, images)
}
