package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	tz := "Asia/Tokyo"
	jst := time.FixedZone(tz, 9*60*60)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")

	r.GET("/", func(c *gin.Context) {
		now := time.Now().In(jst)

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"today":     now.Format("2006/01/02"),
			"dayOfWeek": now.Weekday(),
			"now":       now.Format("15:04:05"),
			"timeZone":  tz,
		})
	})

	r.Run(":8000")
}
