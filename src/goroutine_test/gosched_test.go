package goroutine_test

import (
    "fmt"
    "runtime"
    "testing"
)

func TestGoroutine(t *testing.T) {
    go func() {
        for i := 0; i < 10; i++ {
            go fmt.Println("this is son goroutine")
        }
    }()
    for {
        runtime.Gosched()
        fmt.Println("🍀 this is main")
    }

}
