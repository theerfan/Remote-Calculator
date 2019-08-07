package main

import (
    "fmt"
    "net"
	"os"
	"encoding/gob"
    "github.com/theerfan/Remote-Calculator/util"
    "sync"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type equation = util.Equation

func main() {
    // Listen for incoming connections.
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
    mux := &sync.Mutex{}
    for {
        // Listen for an incoming connection.
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            os.Exit(1)
        }
        // Handle connections in a new goroutine.
        go handleRequest(conn, mux)
    }
}

// Handles incoming requests.
func handleRequest(conn net.Conn, lock *sync.Mutex) {
	var eq equation
	for {
		enc := gob.NewEncoder(conn)
		dec := gob.NewDecoder(conn)
		err := dec.Decode(&eq)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}
		ans := fmt.Sprintf("Answer is %f", calculate(eq))
		fmt.Println(ans)
		lock.Lock()
		error := enc.Encode(ans)
		lock.Unlock()
		if error != nil {
			fmt.Println("Connection closed")
			break
		}
	}	
	conn.Close()
}

