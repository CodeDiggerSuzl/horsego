package goroutine

import (
    "fmt"
    "testing"
    "time"
)

// global chan
var channel = make(chan int)

func Test_channel(t *testing.T) {
    go printHello()
    go printWorld()
    for {

    }
}

func printer(s string) {
    for _, ch := range s {
        fmt.Printf("%c", ch)
        time.Sleep(time.Millisecond * 300)
    }
}

// first
func printHello() {
    printer("hello")
    channel <- 1
}

// second
func printWorld() {
    num := <-channel
    fmt.Println(num)
    printer("world")
}
