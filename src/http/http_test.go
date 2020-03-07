package http

import (
    "fmt"
    "net"
    "os"
    "testing"
)

func HandleErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        // return 返回当前函数调用
        // runtime.Goexit() 结束当前 go 程
        os.Exit(1) // 将当前进程结束
    }
}
func Test_Http(t *testing.T) {
    listener, err := net.Listen("tcp", ":8000")
    HandleErr(err)
    defer listener.Close()

    conn, err := listener.Accept()
    HandleErr(err)
    defer conn.Close()

    buf := make([]byte, 4096)
    n, err := conn.Read(buf)
    HandleErr(err)
    if n == 0 {
        return
    }

    fmt.Printf("🥕%s🥕\n", string(buf[:n]))
}
