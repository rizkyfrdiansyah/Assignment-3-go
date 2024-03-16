package web

import (
	"net/http"
	"weather-monitoring/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupServer(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/weather", func(c *gin.Context) {
		var weather models.Weather
		db.Last(&weather)

		status := map[string]string{
			"water": waterStatus(weather.Water),
			"wind": windStatus(weather.Wind),
		}

		c.JSON(http.StatusOK, gin.H{"status": status})
	})

	r.Static("/static", "./web/static")

	r.LoadHTMLFiles("./web/templates/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return r
}

func waterStatus(value int) string {
	switch {
	case value < 5:
		return "aman"
	case value >= 6 && value <= 8:
		return "siaga"
	default:
		return "bahaya"
	}
}

func windStatus(value int) string {
	switch {
	case value < 6:
		return "aman"
	case value >= 7 && value <= 15:
		return "siaga"
	default:
		return "bahaya"
	}
}