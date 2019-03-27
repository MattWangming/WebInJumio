package netverify

import (
	"testing"
)

func TestInitiate(t *testing.T) {
	Initiate()
}

func TestRetrievaldata(t *testing.T) {
	scanref := "5a3fdc60-a6a3-4a04-aef3-428b7c7d38f0"
	flag := "data"
	result := Retrievaldata(scanref,flag)
	t.Log(result)
}