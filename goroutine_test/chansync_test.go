package goroutine

import (
    "fmt"
    "testing"
    "time"
)

func TestChanSync(t *testing.T) {

    ch := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            fmt.Println("one write", i)
            ch <- i
        }
    }()
    time.Sleep(time.Millisecond * 300)
    go func() {
        for i := 0; i < 5; i++ {
            read := <-ch
            fmt.Println("second read", read)
        }
    }()
}
