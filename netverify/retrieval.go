package netverify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func Retrievaldata(scanReference, flag string) string {
	BaseURL := "https://netverify.com/api/netverify/v2/scans/"
	//assemble the url with different router via different demands, e.g. data, transaction only, .etc
	var url string

	switch{
	//if nothing input, then justify the status of the scan reference
	case flag == "":
		url = BaseURL + scanReference

	//if "data" input, then query all the data information
	case flag == "data":
		url = BaseURL + scanReference + "/" + flag

	//only want document data
	case flag == "document":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	//only want transaction data
	case flag == "transaction":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	//only want verification data
	case flag == "verification":
		url = BaseURL + scanReference + "/" + "data" + "/" + flag

	//only the front images
	case flag == "front":
		url = BaseURL + scanReference + "/" + "images" + "/" + flag

	//only the back images, when ID_CARD is selected as ID type
	case flag == "back":
		url = BaseURL + scanReference + "/" + "images" + "/" + flag

	//only the face images
	case flag == "face":
		url = BaseURL + scanReference + "/" + "images" + "/" + flag


	default:
		return fmt.Sprintf("No such corresponding fields, please input the correct flag!")
	}

	//initiate a http request obj with GET method
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	//application/json or image/jpeg, image/png for "Retrieving specific image"
	req.Header.Set("Accept", "application/json")
	req.Header.Add("Accept", "image/png")
	req.Header.Add("Accept", "image/jpeg")
	//add the user agent for trouble shooting
	req.Header.Add("User-Agent", "Digital Wallet QSTOApp/v1.0")

	//note the specific Authorization code with jumio API credential, updated accordingly!
	req.Header.Add("Authorization","Basic OWJjZmFhM2QtNThkMy00MjhlLWE5ZTUtYzM3YTc4NDZjMjUwOkFLVXppVjFlNGo2WndYQ2d2SDR4d0o2dGlnUVFxc2Fi")


	//get the origin codes from the standard net/http package
	resp, err := http.DefaultClient.Do(req)
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

func RetrievalServer()  {
	r := gin.Default()
	//set the CORS policy
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD, OPTIONS, PUT, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type") //有使用自定义头 需要这个,Action, Module是例子

		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	})
	r.GET("/retrieval", func(c *gin.Context) {
			resp := Retrievaldata("948cc1c2-200e-42be-89c1-bf4113a083d1", "data")
			fmt.Print(string(resp))
			c.JSON(200, string(resp))
	})
	r.Run("192.168.1.23:8849")
}
