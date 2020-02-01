package tcpip

import (
    "fmt"
    "net"
    "os"
    "testing"
)

func Test_UdpClient(t *testing.T) {
    conn, err := net.Dial("udp", ":8848")
    if err != nil {
        fmt.Println("❌ err during udp dial", err)
        return
    }
    defer conn.Close()

    fmt.Println("✅ udp conn created correctly")

    go func() {
        str := make([]byte, 1024)
        for {
            rdStr, err := os.Stdin.Read(str)
            if err != nil {
                fmt.Println("❌ err during os.Stdin")
                return
            }
            conn.Write(str[:rdStr])
        }
    }()
    // conn.Write([]byte("hello from udp client"))

    buf := make([]byte, 4096)
    for {
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("❌ err during conn.Read")
            return
        }
        fmt.Println("read from udp server", string(buf[:n]))
    }
}
