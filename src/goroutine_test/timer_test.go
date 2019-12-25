package goroutine

import (
    "fmt"
    "testing"
    "time"
)

func TestTimer(t *testing.T) {
    fmt.Println(time.Now())
    timer := time.NewTimer(2 * time.Second)
    nowTime := <-timer.C
    fmt.Println(nowTime)
}
func TestTimerAfter(t *testing.T) {
    // 1. sleep
    time.Sleep(time.Second)
    // 2. timer
    fmt.Println(time.Now())
    timer := time.NewTimer(2 * time.Second)
    nowTime := <-timer.C
    fmt.Println(nowTime)
    // 3. time.After
    nowTime = <-time.After(3 * time.Second)
}
