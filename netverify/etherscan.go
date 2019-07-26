package netverify

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Ethinput struct {
	Addr   string `form:"addr"`
	Page   string `form:"page"`
	Offset string `form:"offset"`
}

//Etherscantxlist
func Etherscantxlist() {
	r := gin.Default()
	r.GET("/txlist", func(c *gin.Context) {
		var ethinput Ethinput
		if c.ShouldBind(&ethinput) == nil {
			resp := Ethscan(ethinput.Addr, ethinput.Page, ethinput.Offset)
			c.JSON(200, resp)
		}

	})
	r.Run(":8848")
}

//Ethscan
func Ethscan(addr, page, offset string) string {
	baseurl := "https://api.etherscan.io/api?module=account&action=txlist"
	ad := addr
	pa := page
	os := offset

	url := baseurl + "&address=" + ad + "&startblock=0&endblock=99999999" + "&page=" + pa + "&offset=" + os + "&sort=asc&apikey=YourApiKeyToken"
	// req, _ := http.NewRequest("GET", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(respBody)

}
