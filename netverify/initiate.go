package netverify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)
type preset struct {
	Index 		int    		`json:"index,omitempty"`
	Country		string 		`json:"country,omitempty"`
	Type        string		`json:"type,omitempty"`
	Phrase      string		`json:"phrase,omitempty"`
}

/*  presets template
[
     {
        "index"   : 1,
        "country" : "CHN",
        "type"    : "PASSPORT"
     },
     {
		"index"   : 2,
		"phrase"  : "my custom note"
     }
    ]

 */

 /*  workflowId template
 Value	Verification type	Capture method
 100	ID only				camera + upload
 101	ID only				camera only
 102	ID only				upload only
 200	ID + Identity		camera + upload
 201	ID + Identity		camera only
 202	ID + Identity		upload only

  */

type presets []preset

//send body struct to be delivered to Jumio server
type sendbody2Jumio struct {
	CIR				string 		`json:"customerInternalReference"`
	UserRef			string 		`json:"userReference"`
	SucURL			string		`json:"successUrl"`
	ErrURL			string		`json:"errorUrl"`
	CabURL			string		`json:"callbackUrl"`
	RepCri			string 		`json:"reportingCriteria"`
	WorkfId			int			`json:"workflowId"`
	Presets			presets		`json:"presets"`
	Locale			string		`json:"locale"`
}

var (
	sb sendbody2Jumio
	)

func Post2jumio(url, country, locale, IDtype, presetNote string) []byte {
	//initiate the body with customized part
	timestamp := time.Now().Format("20060102150405")
	sb.CIR = "qsto"
	sb.UserRef = "user" + timestamp
	sb.SucURL = "https://www.qsto.network/marketplace"
	sb.CabURL = "https://www.qsto.network/api/jumio/callback"
	sb.ErrURL = "https://www.qsto.network/marketplace"
	//"https://www.qsto.network/marketplace"
	sb.RepCri = "userReport"
	//per discussion, with 200 default
	sb.WorkfId = 200
	sb.Locale = locale

	//change the preset accordingly
	var preset2Index int
	if presetNote == "" {
		preset2Index = 0
		sb.Presets = presets{
			preset{
				1,
				country,
				IDtype,
				"",
			},
		}
		}else {
		preset2Index = 2
		sb.Presets = presets{
			preset{
				1,
				country,
				IDtype,
				"",
			},
			preset{
				preset2Index,
				"",
				"",
				presetNote,
			},
		}
	}

	//other parts fetch from the request of html
	payload, _ := json.Marshal(sb)
	body := bytes.NewBuffer(payload)
	fmt.Println(string(payload))
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Digital Wallet QSTOApp/v1.0")
	//Note the specific Authorization code with jumio API credential, updated accordingly!
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

	return respBody

}

func Initiate() {
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
	r.POST("/initiate", func(c *gin.Context) {
		//fetch info from request via html, e.g. country, locale, IDtype, presetNote if any, workflowId
		type Formdata struct {
			WorkflowId		int			`json:"workflowId"`
			Country			string		`json:"country"`
			IDType			string		`json:"type"`
			Locale			string		`json:"locale"`
			PresetNote		string      `json:"presetNote,omitempty"`
		}
		var data Formdata

		if c.BindJSON(&data) == nil {
		// url = https://netverify.com/api/v4/initiate
			resp := Post2jumio("https://netverify.com/api/v4/initiate", data.Country, data.Locale, data.IDType, data.PresetNote)
			fmt.Print(string(resp))
			c.JSON(200, string(resp))

		} else {
			c.JSON(404,"no such result")
		}

	})
	r.Run("192.168.1.23:8848")
}

func InitiateMock() {
	r := gin.Default()
	r.POST("/initiate", func(c *gin.Context) {
		randstring := rdString()
		timeStr:= time.Now().Format("2006-01-02T150405")
		c.JSON(200, gin.H{
			"timestamp": timeStr,
			"transactionReference": randstring,
			"redirectUrl": "192.168.1.23:8080/up2",
		})
		//time.Sleep(10 * time.Second)
		//defer RetrievalInfo2Db(randstring)
		//initiate a go routine to post info into DB
		//go func() {
			// simulate a long task with time.Sleep(). 5 minutes
			//time.Sleep(10 * time.Second)
			//RetrievalInfo2Db(randstring)
		//}()
	})
	r.Run("192.168.1.23:8850")
}

func rdString() string {
	const le8 = "886c11b5"
	const le4 = "e9a8"
	const le12  = "b99cdef55e14"

	b8 := make([]byte,8)
	b4 := make([]byte,4)
	b12 := make([]byte, 12)

	for i := range b8 {
		b8[i] = le8[rand.Intn(8) % int(len(le8))]
	}

	for i := range b4 {
		b4[i] = le4[rand.Intn(4) % int(len(le4))]
	}

	for i := range b12 {
		b12[i] = le12[rand.Intn(12) % int(len(le12))]
	}

	randstring := string(b8) + "-" + string(b4) + "-" + string(b4) + "-" + string(b4) + "-" + string(b12)

	return randstring
}