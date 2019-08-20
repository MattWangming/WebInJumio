package netverify

import (
	"github.com/gin-gonic/gin"
)

func Quickscan() {
	r := gin.Default()
	r.GET("/QRdata/ethereum", func(c *gin.Context) {
		//example for input pattern:"cosmos:<address>?contractAddress=uatom&decimal=<decimal>&value=<value>"
		data := `[
					{
						"chain":"^ethereum:$",
						"address":"^0x[0-9a-fA-F]{40}$",
						"contractAddress":"^contractAddress=0x[0-9a-fA-F]{40}$",
						"decimal":"^decimal=[0-18]$",
						"value":"^value=[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0$",	
						"comment":"This is for imToken ethereum template checking"	
					},
					{
						"address":"^0x[0-9a-fA-F]{40}$",
						"comment":"This is for other ethereum template checking"
					}
				]`
		c.String(200, data)
	})
	r.GET("/QRdata/cosmos", func(c *gin.Context) {
		//example for input pattern:"cosmos:<address>?contractAddress=uatom&decimal=<decimal>&value=<value>"
		data := `[
					{	
						"chain":"^cosmos:$",
						"address":"^cosmos1[0-9a-z]{38}$",
						"contractAddress":"^uatom$",
						"decimal":"^decimal=[0-18]$",
						"value":"^value=[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0$",
						"comment":"This is for imtoken cosmos networks template checking"
					},
					{
						"address":"^cosmos1[0-9a-z]{38}$",
						"comment":"This is for other cosmos networks template checking"
					}
					]`
		c.String(200, data)
	})

	r.Run(":8850")
}
