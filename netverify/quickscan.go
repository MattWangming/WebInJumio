package netverify

import (
	"github.com/gin-gonic/gin"
)

func Quickscan() {
	r := gin.Default()
	r.GET("/QRdata", func(c *gin.Context) {
		// if err := c.ShouldBindUri(&jumioRef); err != nil {
		// 	c.JSON(400, gin.H{"msg": err})
		// 	return
		// }
		data := `{
			"ethereum": "<address>",
			"cosmos": "<address>",
			"<address>",
			"contractAddress",
			"decimal",
			"value"
		}`
		c.String(200, data)

	})
	r.Run(":8850")
}
