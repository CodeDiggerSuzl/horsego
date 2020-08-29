package goroutine

import (
    "fmt"
    "runtime"
    "testing"
)

func Test_maxprocess(t *testing.T) {
    gomaxprocs := runtime.GOMAXPROCS(1)
    fmt.Println(gomaxprocs)
    runtime.GC()
    fmt.Println(runtime.NumCPU())
    
    // for {
    //     go fmt.Print(0)
    //     fmt.Print(1)
    // }
}
