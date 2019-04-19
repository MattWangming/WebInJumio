package netverify

import (
	"encoding/json"
	"fmt"
	"github.com/mergemap"
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

type Doc struct {
	Type			string			`json:"type"`
	Dob				string			`json:"dob"`
	Expiry			string			`json:"expire"`
	FirstName		string			`json:"firstName"`
	IssuingCountry  string			`json:"issuingCountry"`
	IssuingDate		string			`json:"issuingDate"`
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

func RetrievalInfo2Db(scanReference string) {
	data := RetrievalfromJumio(scanReference, "data")
	dataBytes := []byte(data)
	img := RetrievalfromJumio(scanReference, "images")
	imgBytes := []byte(img)

	//remove duplicated fields
	var m1, m2 map[string]interface{}
	json.Unmarshal(dataBytes, &m1)
	json.Unmarshal(imgBytes, &m2)
	kycMerg := mergemap.Merge(m1, m2)

	//convert map to byte[]
	kycBz, _ := json.Marshal(kycMerg)

	var KycRes KycResultMerged
	json.Unmarshal(kycBz,&KycRes)

	//info need to push/post to database
	detail := string(kycBz)
	expireTime := KycRes.Document.Expiry
	jumioId := KycRes.ScanReference
	var result int
	//result according to validity
	switch {
	case KycRes.Verification.IdentityVerification.Validity == "true":
		result = 1
	case KycRes.Verification.IdentityVerification.Validity == "false":
		result =2
	default:
		result =0
	}

	fmt.Printf(detail,expireTime,jumioId,result)




}
