package main

import (
    "fmt"
    "net"
    "os"
    "encoding/binary"
    "equation"
    "sync"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

type equation = equation.equation

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
    mux := &Sync.Mutex{}
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
func handleRequest(conn net.Conn, lock *Sync.Mutex) {
  // Make a buffer to hold incoming data.
  byteBuf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  reqLen, err := conn.Read(byteBuf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  eq := equation{}
  buf := &bytes.Buffer{}
  err = binary.Read(buf, binary.bigEndian, &eq)
  ans := fmt.Sprintf(calculate(eq))
  // Send a response back to person contacting us.
  lock.Lock()
  conn.Write([]byte())
  lock.Unlock()
  fmt.Println(eq, ans)
  // Close the connection when you're done with it.
  conn.Close()
}

func calculate(eq equation) float64{
    switch eq.fix {
    case "+":
        return eq.a + eq.b
    case "-":
        return eq.a - eq.b
    case "*":
        return eq.a * eq.b
    case "/":
        return eq.a / eq.b
    }
}