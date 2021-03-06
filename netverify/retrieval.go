package netverify

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mergemap"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func RetrievalfromJumio(scanReference, flag string) string {
	//BaseURL := "https://netverify.com/api/netverify/v2/scans/"
	BaseURL := "http://127.0.0.1:8849/"
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

	//all the images info requery
	case flag =="images":
		url = BaseURL + scanReference + "/" + flag
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
		data := RetrievalfromJumio("948cc1c2-200e-42be-89c1-bf4113a083d1", "data")
		dataBytes := []byte(data)
		img := RetrievalfromJumio("948cc1c2-200e-42be-89c1-bf4113a083d1", "images")
		imgBytes := []byte(img)

		var m1, m2 map[string]interface{}

		json.Unmarshal(dataBytes, &m1)
		json.Unmarshal(imgBytes, &m2)

		kycRes := mergemap.Merge(m1, m2)
		kycBz, _ := json.Marshal(kycRes)

		c.JSON(200, string(kycBz))
	})
	r.Run("192.168.1.23:8849")
}

func RetrievalfromJumioMock() {
	type JumioId struct {
		ScanReference string  `uri:"scanReference" binding:"required"`
	}
	var jumioRef JumioId
	r := gin.Default()
	r.GET("/:scanReference/data", func(c *gin.Context) {
		if err := c.ShouldBindUri(&jumioRef); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		data :=`{
			"timestamp": "2019-03-19T021726",
			"scanReference": "oJnNPGsiuzytMOJPatwtPilfsfykSBGp",
			"document": {
				"type": "PASSPORT",
				"dob": "1985-05-18",
				"expiry": "2022-04-18",
				"firstName": "MING",
				"issuingCountry": "CHN",
				"lastName": "WANG",
				"number": "G61446824",
				"personalNumber": "19201100",
				"status": "APPROVED_VERIFIED"
			},
			"transaction": {
				"clientIp": "114.242.55.129",
				"customerId": "usermatt0318",
				"date": "2019-03-19T02:17:26.499Z",
				"merchantReportingCriteria": "mattReport0318",
				"merchantScanReference": "testmatt0318",
				"source": "WEB_UPLOAD",
				"status": "DONE"
			},
			"verification": {
				"identityVerification": {
					"reason": "SELFIE_IS_SCREEN_PAPER_VIDEO",
					"similarity": "NOT_POSSIBLE",
					"validity": "true"
				},
				"mrzCheck": "OK"
			}
		}`
		dataBytes := []byte(data)
		var datastruct Data
		json.Unmarshal(dataBytes,&datastruct)
		c.Data(200,"application/json",dataBytes)

	})
	r.GET("/:scanReference/images", func(c *gin.Context) {
		if err := c.ShouldBindUri(&jumioRef); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{
			"timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"images": `[
			{
				"classifier": "back",
				"href": "http://aoe-qos.oss-cn-beijing.aliyuncs.com/test/1"
			},
			{
				"classifier": "front",
				"href": "http://aoe-qos.oss-cn-beijing.aliyuncs.com/test/2"
			},
			{
				"classifier": "face",
				"href": "http://aoe-qos.oss-cn-beijing.aliyuncs.com/test/2"
			}
		]`,
			"livenessImages": `[
			"http://aoe-qos.oss-cn-beijing.aliyuncs.com/test/1",
			"http://aoe-qos.oss-cn-beijing.aliyuncs.com/test/2"
		]`,
			"scanReference": jumioRef.ScanReference,
		})
	})
	r.Run(":8849")
}