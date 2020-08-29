package tcpip

import (
    "bytes"
    "fmt"
    "io"
    "net"
    "os"
    "testing"
)

func TestSmpHTTP(t *testing.T) {
    conn, err := net.Dial("tcp", "qbox.me:80")
    handleErr(err)

    _, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
    handleErr(err)

    result, err := readFully(conn)
    handleErr(err)
    fmt.Printf(string(result))

    os.Exit(0)
}

func readFully(conn net.Conn) ([]byte, error) {
    defer conn.Close()

    result := bytes.NewBuffer(nil)
    var buf [512]byte

    for {
        n, err := conn.Read(buf[0:])
        result.Write(buf[0:n])
        if err != nil {
            if err == io.EOF {
                break
            }
            return nil, err
        }
    }
    return result.Bytes(), nil
}
func handleErr(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fital error: %s", err.Error())
        os.Exit(1)
    }
}
