package goroutine

import (
    "fmt"
    "runtime"
    "testing"
    "time"
)

func TestSelect(t *testing.T) {
    quit := make(chan bool)
    intsChan := make(chan int)
    go func() {
        for i := 0; i < 5; i++ {
            intsChan <- i
            time.Sleep(time.Second)
        }
        close(intsChan)
        quit <- true
        runtime.Goexit()
    }()
    for {
        select {
        case <-intsChan:
            num := <-intsChan
            fmt.Println("🍀 get num:", num)
        case <-quit:
            return
        }
        fmt.Println("🍀")
    }
}

// Fibonacci TODO why deadlock ?
func TestFibonacci(t *testing.T) {
    ch := make(chan int)
    quit := make(chan bool)
    x, y := 1, 1
    for i := 0; i < 10; i++ {
        ch <- x
        x, y = y, x+y
    }
    go fibonacci(ch, quit)
    quit <- true

}
func fibonacci(ch <-chan int, quit <-chan bool) {
    for {

        select {
        case num := <-ch:
            fmt.Println(num)
        case <-quit:
            return
        }
    }
}
func TestExpire(t *testing.T) {
    ch := make(chan int)
    // like a flag in java
    quit := make(chan bool)

    go func() {
        for {
            select {
            // if select this print num the for loop one time
            case num := <-ch:
                fmt.Println("num :", num)
            // Re-timing
            case <-time.After(3 * time.Second):
                // after 3 seconds write a bool to quit and goto mark
                quit <- true // true or false is all ok

                // 1. can't use break to jump out of for loop
                // 2. can use return or runtime.goexit()
                // or goto key word: you can goto every, in this case, we goto outside of for loop
                goto mark

            }
        }
    mark:
        fmt.Println("go to here")
    }()
    //
    for i := 0; i < 100; i++ {
        ch <- i
        time.Sleep(3 * time.Second) // write to chan every 3 seconds
    }
    <-quit // wait for goroutine to notice the test-routine
    fmt.Println("all done")
}
