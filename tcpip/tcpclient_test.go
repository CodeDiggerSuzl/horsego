package tcpip

import (
    "fmt"
    "net"
    "testing"
)

func Test_TcpClient(t *testing.T) {
    conn, err := net.Dial("tcp", ":8000")
    if err != nil {
        fmt.Println(" err during dial", err)
    }
    defer conn.Close()
    conn.Write([]byte("R U ready ?"))
    // get from server
    buffer := make([]byte, 4096)
    read, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("err during reading from server")
    }
    fmt.Println("read from server: ", string(buffer[:read]))
}
