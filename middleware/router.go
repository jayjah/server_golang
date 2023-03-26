package middleware

import (
	"github.com/Depado/ginprom"
	"github.com/gin-gonic/gin"
	"golang_server/middleware/routes"
	"net/http"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	p := ginprom.New(
		ginprom.Engine(r),
		ginprom.Subsystem("gin"),
		ginprom.Path("/metrics"),
	)

	r.Use(p.Instrument())
	r.Use(logger)
	r.Use(errorHandler)

	r.StaticFS("/public", http.Dir("public"))

	r.GET("/ping", routes.Ping)

	v4 := r.Group("/v4")
	{
		v4.GET("/events", routes.GetEvents)
		v4.GET("/images", routes.GetImages)
	}

	return r
}
