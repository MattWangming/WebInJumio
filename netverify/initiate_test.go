package netverify

import (
	"encoding/json"
	"fmt"
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


func TestInitiateMock(t *testing.T) {
	InitiateMock()
}

func TestRetrievalfromJumioMock(t *testing.T) {
	RetrievalfromJumioMock()
}

//func TestRetrievalInfo2Db(t *testing.T) {
//	scanReference := "oJnNPGsiuzytMOJPatwtPilfsfykSBGp"
//	RetrievalInfo2Db(scanReference)
//}

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

func TestCombine2(t *testing.T) {
	Combine()
}