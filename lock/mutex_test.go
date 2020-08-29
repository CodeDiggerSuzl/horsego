package lock

import (
    "fmt"
    "sync"
    "testing"
    "time"
)

var mutex sync.Mutex

func TestMutex(t *testing.T) {
    go sayHello()
    go sayWorld()
    for {

    }
}

func printer(str string) {
    mutex.Lock()
    for _, ch := range str {
        fmt.Printf("%c", ch)
        time.Sleep(300 * time.Millisecond)

    }
    mutex.Unlock()
}
func sayHello() {
    printer("hello")
}
func sayWorld() {
    printer("world")
}
