package handlers

import (
	"comete-server/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func GetWeatherByCityName() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var BASE_URL = os.Getenv("BASE_URL")
		var key = os.Getenv("KEY")
		cityName := ctx.Param("city")
		mode := ctx.Param("mode")

		u, err := url.Parse(BASE_URL)
		if err != nil {
			fmt.Println("Erreur lors de l'analyse de l'URL de base:", err)
			return
		}

		q := u.Query()
		q.Set("key", key)
		q.Set("place_id", cityName)
		u.RawQuery = q.Encode()
		newURL := u.String()

		data, err := utils.GetDataFromeApi(newURL)
		if err != nil {
			return
		}
		var m map[string]interface{}
		err = json.Unmarshal([]byte(data), &m)
		if err != nil {
			fmt.Println("Erreur de décodage JSON:", err)
			return
		}

		switch mode {
		case "day":
			ctx.IndentedJSON(http.StatusOK, m["daily"])
		case "hourly":
			ctx.IndentedJSON(http.StatusOK, m["hourly"])
		case "all":
			ctx.IndentedJSON(http.StatusOK, m)
		}
	}
}

func GetWeatherByDay() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var BASE_URL = os.Getenv("BASE_URL")
		var key = os.Getenv("KEY")
		cityName := ctx.Param("city")

		u, err := url.Parse(BASE_URL)
		if err != nil {
			fmt.Println("Erreur lors de l'analyse de l'URL de base:", err)
			return
		}

		q := u.Query()
		q.Set("key", key)
		q.Set("place_id", cityName)
		u.RawQuery = q.Encode()
		newURL := u.String()
		data, err := utils.GetDataFromeApi(newURL)
		if err != nil {
			return
		}
		ctx.IndentedJSON(http.StatusOK, data)
	}
}
