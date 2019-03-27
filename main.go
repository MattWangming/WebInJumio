package main

import (
	"encoding/json"
	"fmt"
	"github.com/WebInJumio/netverify"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Jresp []byte

func main() {
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
			Jresp = netverify.Post2jumio("https://netverify.com/api/v4/initiate", data.Country, data.Locale, data.IDType, data.PresetNote)
			fmt.Print(string(Jresp))
			c.JSON(200, string(Jresp))

		} else {
			c.JSON(404,"no such result")
		}

	})
	r.Run("192.168.1.23:8848")


	//get ready to fetch the data from jumio retrieval API
	type res struct {
		Timestamp		string		`json:"timestamp"`
		TransRef		string		`json:"transactionReference"`
		Redirect		string		`json:"redirectUrl"`
	}

	var rep res
	json.Unmarshal(Jresp,&rep)

	//wait for 4 minutes
	time.Sleep(4 * time.Minute)

	retrivalresult := netverify.Retrievaldata(rep.TransRef, "data")
	fmt.Println(retrivalresult)
}
