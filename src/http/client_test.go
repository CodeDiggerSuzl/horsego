package http

import (
    "fmt"
    "net/http"
    "testing"
)

func Test_Client(t *testing.T) {
    resp, err := http.Get("http://www.baidu.com")
    HandleErr(err)
    defer resp.Body.Close()

    buf := make([]byte, 4096)
    var result string
    for {
        n, _ := resp.Body.Read(buf)
        if n == 0 {
            fmt.Println("Body.Read err:", err)
            break
        }
        result += string(buf[:n])
    }
    fmt.Println(result)
}
