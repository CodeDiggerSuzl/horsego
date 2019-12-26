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

// timer stop
func TestTimerStop(t *testing.T) {
    // ⏰
    timer := time.NewTimer(time.Second * 2)
    timer.Reset(time.Second)
    go func() {
        <-timer.C
        fmt.Println("go func timer stop")
    }()
    // timer.Stop()
    for {

    }
}

// ticker
func TestTicker(t *testing.T) {
    // flag to to stop the goroutine
    quit := make(chan bool)
    ticker := time.NewTicker(time.Second)
    i := 0
    go func() {
        for {
            nowTime := <-ticker.C
            i++
            fmt.Println("nowTime:", nowTime)
            if i == 8 {
                quit <- true //
            }
        }
    }()
    <-quit // 读取 quit 中的数据 阻塞 test 的 goroutine
}
