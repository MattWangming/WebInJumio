package netverify

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)
type preset struct {
	Index 		int    		`json:"index"`
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

var sb sendbody2Jumio

func post2jumio(url, country, locale, IDtype, presetNote string, workflowId int) []byte {
	//initiate the body with customized part
	var nonce int64
	nonce = int64(0)
	nonce++

	sb.CIR = "qsto"
	sb.UserRef = "user" + string(nonce)
	sb.SucURL = "https://www.qsto.network/api/jumio/success"
	sb.CabURL = "https://www.qsto.network/api/jumio/callback"
	sb.ErrURL = "https://www.qsto.network/api/jumio/error"
	sb.RepCri = "userReport"
	sb.WorkfId = workflowId
	sb.Locale = locale

	//change the preset accordingly
	var preset2Index int
	if presetNote == "" {
		preset2Index = int(nil)
	} else {
		preset2Index = 2
	}
	sb.Presets = presets{
		preset{
			1,
			country,
			IDtype,
			nil,
		},
		preset{
			preset2Index,
			nil,
			nil,
			presetNote,
		},
	}
	//other parts fetch from the request of html
	payload, _ := json.Marshal(sb)
	body := bytes.NewBuffer(payload)

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Digital Wallet QSTOApp/v1.0")
	//Note the specific Authorization code with jumio API credential, updated accordingly!
	req.Header.Add("Authorization","Basic OWJjZmFhM2QtNThkMy00MjhlLWE5ZTUtYzM3YTc4NDZjMjUwOkFLVXppVjFlNGo2WndYQ2d2SDR4d0o2dGlnUVFxc2Fi")
	clt := http.Client{}
	resp, _ := clt.Do(req)
	defer resp.Body.Close()
	rep, _ := ioutil.ReadAll(resp.Body)
	return rep
}

func Initiate() {
	r := gin.Default()
	r.POST("/initiate", func(c *gin.Context) {
		//fetch info from request via html, e.g. locale, presets:country,type, workflowId designed



		//post http request to jumio api: https://netverify.com/api/v4/initiate
		resp := post2jumio("https://netverify.com/api/v4/initiate")
		c.JSON(200, resp)
	})
	r.Run("localhost:8081")
}