package tcpip

import (
    "fmt"
    "net"
    "runtime"
    "strings"
    "testing"
)

// Test Concurrent Server
func Test_CoTcpServer(t *testing.T) {
    // create listen socket
    listener, err := net.Listen("tcp", ":8000")
    if err != nil {
        fmt.Println("net.Listener err", err)
        return
    }
    defer listener.Close()
    for {
        // create conn
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("conn err", err)
            return
        }
        defer conn.Close()
        // a func to handle client and server connection
        go HandlerConnect(conn)
    }

}

func HandlerConnect(conn net.Conn) {
    // get system os and find out use \r\n or only \n
    sysType := runtime.GOOS
    fmt.Println(sysType)
    sysSep := 1
    if strings.Contains(sysType, "windows") {
        sysSep = 2
    }
    defer conn.Close()
    // get client addr
    addr := conn.RemoteAddr()
    fmt.Println(addr, "connect correctly")
    buffer := make([]byte, 4096)
    conn.Write([]byte("connected you can send now ~ \n"))
    for {
        fmt.Println("wait to read from client")
        n, err := conn.Read(buffer)

        if n == 0 {
            fmt.Println("server detected the client was closed")
            return
        }
        if err != nil {
            fmt.Println(err, "during reading ")
            return
        }
        strFromClient := string(buffer[:n])
        fmt.Println("server read from date", strFromClient)
        fmt.Println(buffer[:n])
        // convert string to uppercase
        if "exit" == string(buffer[:n-sysSep]) { // 自己写的客户端测试, 发送时，多了2个字符, "\r\n"
            // if "exit\n" == string(buffer[:n-sysSep]) { // 自己写的客户端测试, 发送时，多了2个字符, "\r\n"
            fmt.Println(addr, "exit")
            conn.Write([]byte("bye \n"))
            return
        }
        conn.Write([]byte(strings.ToUpper(strFromClient)))
    }
}
