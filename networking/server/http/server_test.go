package http

import (
	"testing"
)

func BenchmarkCalculateHandler(b *testing.B){
	jsondata := "{\"A\":2,\"B\":2,\"Fix\":\"+\"}"
	jsonbyte := []byte(jsondata)
	for n := 0; n < b.N; n++ {
		CalculateHandler(jsonbyte)
	}
}