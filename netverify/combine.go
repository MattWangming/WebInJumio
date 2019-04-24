package netverify

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func Combine() {
	r := gin.Default()
	r.POST("/initiate", func(c *gin.Context) {
		const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
		b := make([]byte, 32)
		for i := range b {
			b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
		}
		randstring := string(b)
		timeStr := time.Now().Format("2006-01-02T150405")
		c.JSON(200, gin.H{
			"timestamp":            timeStr,
			"transactionReference": randstring,
			"redirectUrl":          "www.abc.com",
		})

		ch := make(chan string, 1000)
		ch <- randstring
		go RetrievalInfo2Db(ch)
	})
	r.Run("192.168.1.23:8850")
}
