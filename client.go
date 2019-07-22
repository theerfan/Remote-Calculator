package main
import (
    "net"
	"os"
	"strings"
	"bufio"
	"fmt"
	"strconv"
	"equation"
)

type equation = equation.equation

func sendRequest(conn net.Conn, eq equation){
	_, err = conn.Write([]byte(eq))
    if err != nil {
        fmt.Println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    // println("write to server = ", strEcho)

    reply := make([]byte, 1024)
    _, err = conn.Read(reply)
    if err != nil {
        fmt.Println("Write to server failed:", err.Error())
        os.Exit(1)
    }

    fmt.Println("reply from server=", string(reply))

    conn.Close()
}

func extractEquation(input string) equation {
	strArr := strings.Split(input, " ")
	a, err = strconv.ParseFloat(strArr[0])
	fix = strArr[1]
	b, err = strconv.ParseFloat(strArr[2])
	eq := equation{a: a, fix: fix, b: b}
	return eq
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
	defer conn.Close()    

	reader := bufio.NewReader(os.stdin)
	input := reader.ReadString("\n")
	for input != "quit" {
		eq := extractEquation(input)
		go sendRequest()
	} 
}