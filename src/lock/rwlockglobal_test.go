package lock

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

// 使用全局变量 模拟共享数据

var rwMutexG sync.RWMutex
var globalVar int

func Test_RwLockGlobal(t *testing.T) {
    // random seed
    rand.Seed(time.Now().UnixNano())
    quit := make(chan bool)
    for i := 0; i < 5; i++ {
        go readGoG(i + 1)
    }
    for i := 0; i < 5; i++ {
        go writeGoG(i + 1)
    }
    // keep running
    <-quit
}

func readGoG(idx int) {
    for {
        rwMutexG.RLock() // 这样写 得到解锁
        // 🔐 的粒度问题 readLock
        num := globalVar
        fmt.Printf("----%d read %d\n", idx, num)
        rwMutexG.RUnlock()
        time.Sleep(time.Second)
    }
}

func writeGoG(idx int) {
    rwMutexG.Lock()
    for {
        num := rand.Intn(1000)
        rwMutexG.Lock()
        globalVar = num
        fmt.Printf("%d write %d\n", idx, num)
        time.Sleep(time.Millisecond * 400)
        rwMutexG.Unlock()
    }
}
