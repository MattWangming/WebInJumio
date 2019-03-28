package netverify

import (
	"testing"
)

func TestInitiate(t *testing.T) {
	Initiate()
}

func TestRetrievaldata(t *testing.T) {
	scanref := "948cc1c2-200e-42be-89c1-bf4113a083d1"
	flag := "face"
	result := Retrievaldata(scanref,flag)
	t.Log(result)
}