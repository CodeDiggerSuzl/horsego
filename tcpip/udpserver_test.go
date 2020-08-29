package tcpip

import (
    "fmt"
    "net"
    "testing"
    "time"
)

func Test_UdpServer(t *testing.T) {
    // 组织地址结构,制定服务器的 ip port
    addr, err := net.ResolveUDPAddr("udp", ":8848")
    if err != nil {
        fmt.Println("❌ err during ResolveUdpAddr")
        return
    }
    fmt.Println("✅ 服务器地址创建完成")
    // 创建用于通信的 socket
    conn, err := net.ListenUDP("udp", addr)
    if err != nil {
        fmt.Println("❌ err during ListenUDP")
    }
    fmt.Println("✅ 服务器地址链接创建完成")

    defer conn.Close()
    for {

        buf := make([]byte, 2048)
        n, cltAddr, err := conn.ReadFromUDP(buf)
        if err != nil {
            fmt.Println("❌ read from udp err")
            return
        }
        mesFromClt := string(buf[:n])
        fmt.Printf("read from client which is %v:", cltAddr)
        fmt.Println("read from client ", mesFromClt)

        daytime := time.Now().String()
        _, err = conn.WriteToUDP([]byte(daytime), cltAddr)
        if err != nil {
            fmt.Println("❌ err during write 2 client")
            return
        }
    }
}
