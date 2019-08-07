package http

import (
	"testing"
	"net/http"
	"encoding/json"
)

var eq  = equation{
	A: 2,
	B: 3,
	Fix: "+",
}

func BechmarkSendRequest(b *testing.B) {
	client := &http.Client{}
	data, err := json.Marshal(eq)
	if err != nil {
		b.Error("Problematic!" + err.Error())
	}
	for n := 0; n < b.N; n++ {
		sendRequest(client, data)
	}
}