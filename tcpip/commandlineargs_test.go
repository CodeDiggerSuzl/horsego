package tcpip

import (
    "fmt"
    "os"
    "testing"
)

func Test_CommandLineArgs(t *testing.T) {
    // get command line args
    argList := os.Args
    fmt.Println(argList)
}
