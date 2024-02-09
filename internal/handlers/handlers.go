package handlers

import (
	"comete-server/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherByCityName() gin.HandlerFunc {
	url := "https://www.meteosource.com/api/v1/free/point?place_id=antananarivo&sections=all&timezone=UTC&language=en&units=metric&key=0uur5oqw6oixnlah2insijvgtl48mr7v7uaxfbm3"
	return func(ctx *gin.Context) {
		data, err := utils.GetDataFromeApi(url)
		if err != nil {
			return
		}
		ctx.IndentedJSON(http.StatusOK, string(data))
	}
}

func GetWeatherByDay() gin.HandlerFunc {
	url := "https://www.meteosource.com/api/v1/free/"
	return func(ctx *gin.Context) {
		data, err := utils.GetDataFromeApi(url)
		if err != nil {
			return
		}
		ctx.IndentedJSON(http.StatusOK, data)
	}
}
