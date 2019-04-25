package netverify

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

func TestJumIdrefa(t *testing.T) {
	//const letterBytes = "775c11b4-e9a7-4a23-bff8-b99cdef55e13"
	const le8 = "775c11b4"
	const le4 = "e9a7"
	const le12  = "b99cdef55e13"

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

	//b := make([]byte, 32)
	//for i := range b {
	//	b[i] = letterBytes[rand.Int31() % int32(len(letterBytes))]
	//}
	//randstring := string(b)
	t.Log(randstring)
}