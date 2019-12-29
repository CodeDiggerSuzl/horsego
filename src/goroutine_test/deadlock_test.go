package goroutine

import (
    "fmt"
    "testing"
)

// chan must be used in tow more go routines
func TestDLSingleGoroutine(t *testing.T) {
    ch := make(chan int)
    ch <- 9
    num := <-ch
    fmt.Println(num)
}

// on side read/write another must have chance to write/read
func TestDL(t *testing.T) {
    ch := make(chan int)

    num := <-ch
    fmt.Println(num)
    go func() {
        ch <- 32
    }()
}
func TestDLMultiGoroutine(t *testing.T) {
    ch1 := make(chan int)
    ch2 := make(chan int)
    go func() {
        for {
            select {
            case num := <-ch1:
                ch2 <- num
            }
        }
    }()

    for {
        select {
        case num := <-ch2:
            ch1 <- num
        }
    }
}
