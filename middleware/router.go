package middleware

import (
	"github.com/gin-gonic/gin"
	"golang_server/middleware/routes"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()

	r.Use(logger)
	r.Use(errorHandler)

	v4 := r.Group("/v4")
	{
		v4.GET("/events", routes.GetEvents)
	}

	/*r.GET("/events", func(c *gin.Context) {

	})*/

	return r
}
