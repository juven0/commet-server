package routes

import (
	"comete-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	r := gin.New()
	r.GET("/api/comete/:city/:mode", handlers.GetWeatherByCityName())
	return r
}
