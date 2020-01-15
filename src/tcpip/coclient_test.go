package tcpip

import (
    "fmt"
    "net"
    "os"
    "testing"
)

// Test Concurrent Tcp Client

func Test_TcpCoClinet(t *testing.T) {
    conn, err := net.Dial("tcp", ":8000")
    if err != nil {
        fmt.Println("err during net.Dial", err)
        return
    }
    fmt.Println("connect success to server ~")
    defer conn.Close()
    go func() {
        buffer := make([]byte, 4096)
        for {
            n, err := os.Stdin.Read(buffer)
            if err != nil {
                fmt.Println("err during Stdin.Read err", err)
                continue
            }
            _, err = conn.Write(buffer[:n])
            if err != nil {
                fmt.Println("conn.Write err", err)
                return
            }
        }
    }()

    // get from server
    sWriterBuf := make([]byte, 1024)
    for {
        n, err := conn.Read(sWriterBuf)
        if err != nil {
            fmt.Println("conn.Read err", err)
            return
        }
        fmt.Printf("get from server %s\n", string(sWriterBuf[:n]))
    }
}
