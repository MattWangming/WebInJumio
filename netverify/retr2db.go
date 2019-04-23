package netverify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mergemap"
	"io/ioutil"
	"log"
	"net/http"
)
//data structure according to Jumio implemetation guide line, to be updated with normal Jumio license!
type KycResultMerged struct {
	Timestamp			string		`json:"timestamp"`
	ScanReference		string		`json:"scanReference"`
	Document			Doc			`json:"document"`
	Transaction			Tx			`json:"transaction"`
	Verification		Veri      	`json:"verification"`
	Images				Imgs		`json:"images"`
	LivenessImages		[]string	`json:"livenessImages"`
}

type Data struct {
	Timestamp			string		`json:"timestamp"`
	ScanReference		string		`json:"scanReference"`
	Document			Doc			`json:"document"`
	Transaction			Tx			`json:"transaction"`
	Verification		Veri      	`json:"verification"`
}

type Imgages struct {
	Timestamp			string		`json:"timestamp"`
	ScanReference		string		`json:"scanReference"`
	Images				Imgs		`json:"images"`
	LivenessImages		[]string	`json:"livenessImages"`
}

type Doc struct {
	Type			string			`json:"type"`
	Dob				string			`json:"dob"`
	Expiry			string			`json:"expiry"`
	FirstName		string			`json:"firstName"`
	IssuingCountry  string			`json:"issuingCountry"`
	IssuingDate		string			`json:"issuingDate,omitempty"`
	LastName		string			`json:"lastName"`
	Number			string			`json:"number"`
	PersonalNumber  string			`json:"personalNumber"`
	Status			string			`json:"status"`
}

type Tx struct {
	ClientIp						string			`json:"clientIp"`
	CustomerId						string			`json:"customerId"`
	Date							string			`json:"date"`
	MerchantReportingCriteria		string			`json:"merchantReportingCriteria"`
	MerchantScanReference			string			`json:"merchantScanReference"`
	Source							string			`json:"source"`
	Status							string			`json:"status"`
}

type Veri struct {
	IdentityVerification			IdentityVeri		`json:"identityVerification"`
	MrzCheck						string				`json:"mrzCheck"`
	RejectReason					RejectR				`json:"rejectReason,omitempty"`
}

type IdentityVeri struct {
	Reason						string				`json:"reason"`
	Similarity					string				`json:"similarity"`
	Validity					string				`json:"validity"`
	HandwrittenNoteMatches		string				`json:"handwrittenNoteMatches,omitempty"`
}

type RejectR struct {
	RejectReasonCode     		string			`json:"rejectReasonCode,omitempty"`
	RejectReasonDescription		string			`json:"rejectReasonDescription,omitempty"`
	RejectReasonDetails       	RejectD			`json:"rejectReasonDetails,omitempty"`
}

type RejectD struct {
	DetailsCode			string			`json:"detailsCode,omitempty"`
	DetailsDescription  string			`json:"detailsDescription,omitempty"`
}

type img struct {
	Classifier			string		`json:"classifier"`
	Href				string		`json:"href"`
}

type Imgs []img

type DataNew struct {
	Document struct {
		Dob            string `json:"dob"`
		Expiry         string `json:"expiry"`
		FirstName      string `json:"firstName"`
		IssuingCountry string `json:"issuingCountry"`
		LastName       string `json:"lastName"`
		Number         string `json:"number"`
		PersonalNumber string `json:"personalNumber"`
		Status         string `json:"status"`
		Type           string `json:"type"`
	} `json:"document"`
	ScanReference string `json:"scanReference"`
	Timestamp     string `json:"timestamp"`
	Transaction   struct {
		ClientIP                  string `json:"clientIp"`
		CustomerID                string `json:"customerId"`
		Date                      string `json:"date"`
		MerchantReportingCriteria string `json:"merchantReportingCriteria"`
		MerchantScanReference     string `json:"merchantScanReference"`
		Source                    string `json:"source"`
		Status                    string `json:"status"`
	} `json:"transaction"`
	Verification struct {
		IdentityVerification struct {
			Reason     string `json:"reason"`
			Similarity string `json:"similarity"`
			Validity   string `json:"validity"`
		} `json:"identityVerification"`
		MrzCheck string `json:"mrzCheck"`
	} `json:"verification"`
}


//complete the flow with Retrieval-->Post2DB
func RetrievalInfo2Db(scanReference string) {
	//retrieval info from Jumio server with the verification results
	data := RetrievalfromJumio(scanReference, "data")
	dataBytes := []byte(data)
	var datastruct Data
	json.Unmarshal(dataBytes,&datastruct)
	fmt.Println(datastruct.Document)

	img := RetrievalfromJumio(scanReference, "images")
	imgBytes := []byte(img)

	//remove duplicated fields
	var m1, m2 map[string]interface{}
	json.Unmarshal(dataBytes, &m1)
	json.Unmarshal(imgBytes, &m2)
	kycMerg := mergemap.Merge(m1, m2)

	//convert map to byte[]
	kycBz, _ := json.Marshal(kycMerg)
	//fmt.Println(string(kycBz))
	var KycRes KycResultMerged
	json.Unmarshal(kycBz,&KycRes)
	jumioId := KycRes.ScanReference
	//info need to push/post to database
	detail := string(kycBz)

	//unmarshal into depth
	docum := KycRes.Document
	fmt.Println(docum)
	//var Docres Doc
	//json.Unmarshal(docum, &Docres)
	expireTime := KycRes.Document.Expiry

	var result int32
	//result according to validity
	switch {
	case KycRes.Verification.IdentityVerification.Validity == "true":
		result = 1
	case KycRes.Verification.IdentityVerification.Validity == "false":
		result = 2
	default:
		result = 0
	}

	fmt.Printf(detail,expireTime,jumioId,result)

	type formData struct {
		Detail				string		`json:"detail"`
		ExpireTime			string		`json:"expireTime"`
		JumioId				string		`json:"jumioId"`
		Result				int32		`json:"result"`
		Cipher				string		`json:"cipher"`
	}

	form := formData{
		detail,
		expireTime,
		jumioId,
		result,
		"123456",
	}
	//set the default callback url
	url := "http://192.168.1.230:9090/kycApplies/callback"

	payload, _ := json.Marshal(form)
	body := bytes.NewBuffer(payload)
	fmt.Println(string(payload))
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(respBody))

}
