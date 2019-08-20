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
						"address":"^0x[0-9a-fA-F]{40}$",
						"comment":"This is for normal ethereum template checking"
					},			
					{
						"address":"^ethereum:0x[0-9a-fA-F]{40}\b",
						"contractAddress":"\bcontractAddress=0x[0-9a-fA-F]{40}\b",
						"decimal":"\bdecimal=\d+(\.\d+)?\b",
						"value":"\bvalue=\d+(\.\d+)?$",	
						"comment":"This is for imToken ethereum template checking"	
					}
				]`
		c.String(200, data)
	})
	r.GET("/QRdata/cosmos", func(c *gin.Context) {
		//example for input pattern:"cosmos:<address>?contractAddress=uatom&decimal=<decimal>&value=<value>"
		data := `[
					{
						"address":"^cosmos1[0-9a-z]{38}$",
						"comment":"This is for normal cosmos networks template checking"
					},
					{	
						"address":"^cosmos:cosmos1[0-9a-z]{38}\b",
						"contractAddress":"\bcontractAddress=uatom\b",
						"decimal":"\bdecimal=\d+(\.\d+)?\b",
						"value":"\bvalue=\d+(\.\d+)?$",
						"comment":"This is for imtoken cosmos networks template checking"
					}
				]`
		c.String(200, data)
	})

	r.Run(":8850")
}
