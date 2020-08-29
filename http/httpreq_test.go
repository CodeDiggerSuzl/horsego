package http

import (
    "fmt"
    "net"
    "testing"
)

func Test_httpReq(t *testing.T) {
    conn, err := net.Dial("tcp", ":8000")
    HandleErr(err)
    defer conn.Close()
    request := "GET /test HTTP/1.1\r\nHost:127.0.0.1:8000\r\n\r\n"
    conn.Write([]byte(request))
    buf := make([]byte, 4096)
    n, _ := conn.Read(buf)
    if n == 0 {
        return
    }
    fmt.Println(string(buf[:n]))
}
