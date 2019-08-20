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
		//example for input pattern:"cosmos:<address>?contractAddress=uatom&decimal=<decimal>&value=<value>"
		data := `{
			"1" : {
				"ethereum":{
					"address":"^0x[0-9a-fA-F]{40}\?$",
					"contractAddress":"^0x[0-9a-fA-F]{40}\&$",
					"decimal":"^[0-18]\&$",
					"value":"^[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0$"
				},
				"cosmos":{
					"address":"^cosmos1[0-9a-z]{38}\?$",
					"contractAddress":"^uatom\&$",
					"decimal":"^[0-18]\&$",
					"value":"^[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0$"
				}
			},
			"2" : {
				"EthAddress": "^0x[0-9a-fA-F]{40}$",
				"CosmosAddress": "^cosmos1[0-9a-z]{38}$",
				"QosAddress": "address1[0-9a-z]{38}$"
			}
		}`
		c.String(200, data)

	})
	r.Run(":8850")
}
