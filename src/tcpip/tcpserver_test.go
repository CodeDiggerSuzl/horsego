package tcpip

import (
    "fmt"
    "net"
    "testing"
)

// a test to test tcp server
func Test_TcpServer(t *testing.T) {
    listener, err := net.Listen("tcp", ":8000")
    if err != nil {
        fmt.Println("s%", err)
        return
    }
    fmt.Println("ok,set up socket")
    defer listener.Close()
    // block until client connect to server
    conn, err := listener.Accept()

    if err != nil {
        fmt.Println("err during conn :", err)
        return
    }
    // ready to read
    fmt.Println("ready to read from read from client")
    buffer := make([]byte, 4096)
    read, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("err during reading bytes", err)
    }
    fmt.Println("read from clients: ", string(buffer[:read]))
    conn.Write([]byte("got message"))
}
