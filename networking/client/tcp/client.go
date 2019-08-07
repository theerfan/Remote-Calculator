package main

import (
    "net"
	"os"
	"strings"
	"bufio"
	"fmt"
	"strconv"
	"github.com/theerfan/Remote-Calculator/util"
	"encoding/gob"
)

type equation = util.Equation

func sendRequest(conn net.Conn, eq equation){
	enc := gob.NewEncoder(conn)
	dec := gob.NewDecoder(conn)

	err := enc.Encode(eq)
    if err != nil {
        fmt.Println("Write to server failed:", err.Error())
        os.Exit(1)
    }
    fmt.Println("wrote to server = ", eq)

	var reply string
	err = dec.Decode(&reply)
    if err != nil {
        fmt.Println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    fmt.Println("reply from server=", reply)

}

func extractEquation(input string) equation {
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
	return eq
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
    servAddr := "localhost:3333"
    tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
    if err != nil {
        println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
    }

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    if err != nil {
        println("Dial failed:", err.Error())
        os.Exit(1)
	}
	fmt.Println("Successfully connected to " + servAddr)

	reader := bufio.NewReader(os.Stdin)
	input := getInput(reader)
	for input != "quit" {
		eq := extractEquation(input)
		go sendRequest(conn, eq)
		input = getInput(reader)
	} 
	conn.Close()
}