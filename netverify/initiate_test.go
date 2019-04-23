package netverify

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mergemap"
	"testing"
)

func TestInitiate(t *testing.T) {
	Initiate()
}

func TestRetrievaldata(t *testing.T) {
	scanref := "948cc1c2-200e-42be-89c1-bf4113a083d1"
	flag := "data"
	result := RetrievalfromJumio(scanref,flag)
	t.Log(result)
}

func TestRetrievalServer(t *testing.T) {
	RetrievalServer()
}

func TestAppend(t *testing.T)  {
	dataString := `{
	"timestamp": "2019-03-27T04:17:30.795Z",
		"scanReference": "775c11b4-e9a7-4a23-bff8-b99cdef55e13",
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
				"validity": "false"
		},
		"mrzCheck": "OK"
	}
}`
	imgString := `{
   "timestamp": "2019-03-01T11:53:52.878Z",
   "images": [
       {
           "classifier": "back",
           "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/back"
       },
       {
           "classifier": "front",
           "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/front"
       },
       {
           "classifier": "face",
           "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/face"
       }
   ],
		"livenessImages": [
	        	"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/7",
       		"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/5",
       		"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/6",
       		"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/3",
       		"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/4",
		        "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/1",
       		"https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/liveness/2"
   ],
   "scanReference": "775c11b4-e9a7-4a23-bff8-b99cdef55e13"
	}
`
   dataBytes := []byte(dataString)
   imgBytes := []byte(imgString)

	//buf1 := []byte(`{"a":1,"b":2}`)
	//buf2 := []byte(`{"c":3,"d":4,"a":"aaa", "images": [
    //    {
    //        "classifier": "back",
    //        "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/back"
    //    },
    //    {
    //        "classifier": "front",
    //        "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/front"
    //    },
    //    {
    //        "classifier": "face",
    //        "href": "https://netverify.com/api/netverify/v2/scans/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/images/face"
    //    }
    //]}`)

	var m1, m2 map[string]interface{}

    json.Unmarshal(dataBytes, &m1)
    json.Unmarshal(imgBytes, &m2)

    kycRes := mergemap.Merge(m1, m2)
    kycBz, _ := json.Marshal(kycRes)

	t.Log(string(kycBz))
}

func TestInitiateMock(t *testing.T) {
	InitiateMock()
}

func TestRetrievalfromJumioMock(t *testing.T) {
	RetrievalfromJumioMock()
}

func TestRetrievalInfo2Db(t *testing.T) {
	scanReference := "775c11b4-e9a7-4a23-bff8-b99cdef55e13"
	RetrievalInfo2Db(scanReference)
}

func TestJsobStruct(t *testing.T) {
	dataString := `{
	"timestamp": "2019-03-27T04:17:30.795Z",
		"scanReference": "775c11b4-e9a7-4a23-bff8-b99cdef55e13",
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
				"validity": "false"
		},
		"mrzCheck": "OK"
	}
}`
	fmt.Printf("%+v\n",dataString)
	dataBytes := []byte(dataString)
	var datastruct Data
	json.Unmarshal(dataBytes,&datastruct)
	t.Logf("%+v",datastruct.Document)
}

func TestJumRes(t *testing.T)  {
	scanReference:= "775c11b4-e9a7-4a23-bff8-b99cdef55e13"
	data := RetrievalfromJumio(scanReference, "data")
	fmt.Printf("%+v\n",data)
	dataBytes := []byte(data)
	var datastruct Data
	err := json.Unmarshal(dataBytes,&datastruct)
	if err != nil {
		fmt.Println(err)
	}
	t.Log(datastruct.Document)
}

func TestPurejson(t *testing.T)  {
	r := gin.Default()

	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")

}