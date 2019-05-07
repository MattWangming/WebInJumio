package netverify

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Combine() {
	r := gin.Default()
	r.POST("/initiate", func(c *gin.Context) {
		randstring := rdString()
		timeStr := time.Now().Format("2006-01-02T150405")
		c.JSON(200, gin.H{
			"timestamp":            timeStr,
			"transactionReference": randstring,
			"redirectUrl":          "http://0.0.0.0:8080/up2",
		})

		ch := make(chan string, 1000)
		ch <- randstring
		go RetrievalInfo2Db(ch)
	})
	r.Run(":8848")
}
