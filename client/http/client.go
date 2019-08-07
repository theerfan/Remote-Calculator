package main

import (
    "net/http"
	"bytes"
	"strings"
	"bufio"
	"fmt"
	"strconv"
	"github.com/theerfan/Remote-Calculator/util"
	"encoding/json"
	"io/ioutil"
	"os"
)

type equation = util.Equation
const servAddr = "http://127.0.0.1:3333/calculate"

func sendRequest(client *http.Client, body []byte) {
	req, err := http.NewRequest("POST", servAddr, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		fmt.Println("inja")
	}
    // req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    resBody, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(resBody))
}

func extractEquation(input string) []byte {
	strArr := strings.Split(input, " ")
	a, err := strconv.ParseFloat(strArr[0], 64)
	if err != nil{
		fmt.Println("error converting " +  err.Error())
	}
	fix := strArr[1]
	b, err := strconv.ParseFloat(strArr[2], 64)
	if err != nil{
		fmt.Println("error converting " +  err.Error())
	}
	eq := equation{A: a, Fix: fix, B: b}
	data, err := json.Marshal(&eq)
	return data
}

func getInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from input: " + err.Error())
	}
	input = strings.TrimSpace(input)
	return input
}

func main() {
	client := &http.Client{}
	reader := bufio.NewReader(os.Stdin)
	input := getInput(reader)
	for input != "quit" {
		eq := extractEquation(input)
		go sendRequest(client, eq)
		input = getInput(reader)
	} 
}