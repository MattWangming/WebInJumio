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

type Ercinput struct {
	Contract string `form:"contract"`
	Addr     string `form:"addr"`
	Page     string `form:"page"`
	Offset   string `form:"offset"`
}

type Erctoken struct {
	Token string `form:"token"`
}

//Etherscantxlist
func Etherscantxlist() {
	r := gin.Default()
	r.GET("/ethtxlist", func(c *gin.Context) {
		var ethinput Ethinput
		if c.ShouldBind(&ethinput) == nil {
			resp := Ethscan(ethinput.Addr, ethinput.Page, ethinput.Offset)
			c.String(200, resp)
		}

	})
	r.GET("/tokentxlist", func(c *gin.Context) {
		var ercinput Ercinput
		if c.ShouldBind(&ercinput) == nil {
			resp := Ercscan(ercinput.Contract, ercinput.Addr, ercinput.Page, ercinput.Offset)
			c.String(200, resp)
		}

	})
	r.GET("/erctoken", func(c *gin.Context) {
		var erctoken Erctoken
		if c.ShouldBind(&erctoken) == nil {
			resp := Ercsearch(erctoken.Token)
			c.String(200, resp)
		}
	})
	r.Run(":8848")
}

//Ercscantxlist
func Ercscantxlist() {
	r := gin.Default()
	r.GET("/txlist", func(c *gin.Context) {
		var ercinput Ercinput
		if c.ShouldBind(&ercinput) == nil {
			resp := Ercscan(ercinput.Contract, ercinput.Addr, ercinput.Page, ercinput.Offset)
			c.JSON(200, resp)
		}

	})
	r.Run(":8849")
}

//Ethscan
func Ethscan(addr, page, offset string) string {
	baseurl := "https://api.etherscan.io/api?module=account&action=txlist"
	ad := addr
	pa := page
	os := offset

	url := baseurl + "&address=" + ad + "&startblock=0&endblock=99999999" + "&page=" + pa + "&offset=" + os + "&sort=desc&apikey=YourApiKeyToken"
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

//Ercscan
func Ercscan(contract, addr, page, offset string) string {
	baseurl := "https://api.etherscan.io/api?module=account&action=tokentx"
	ca := contract
	ad := addr
	pa := page
	os := offset

	url := baseurl + "&contractaddress=" + ca + "&address=" + ad + "&page=" + pa + "&offset=" + os + "&sort=desc&apikey=YourApiKeyToken"
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

//Ercsearch
func Ercsearch(token string) string {
	baseurl := "https://etherscan.io/searchHandler"
	url := baseurl + "?term=" + token + "&filterby=0"
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
