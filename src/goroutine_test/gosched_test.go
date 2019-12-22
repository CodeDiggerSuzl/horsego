package Goroutine

import (
    "fmt"
    "testing"
    "time"
)

func TestGoroutine(t *testing.T) {
    go func() {
        for i := 0; i < 10; i++ {
            go fmt.Println("this is son goroutine")
            time.Sleep(100 * time.Millisecond)
        }
    }()
    for {
        fmt.Println("🍀 this is main")
        time.Sleep(100 * time.Millisecond)
    }

}
