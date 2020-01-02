package lock

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

// 锁的粒度问题
var rwMutex sync.RWMutex

func Test_RwLock(t *testing.T) {
    // random seed
    rand.Seed(time.Now().UnixNano())
    quit := make(chan bool)
    ch := make(chan int)
    for i := 0; i < 5; i++ {
        go readGo(ch, i+1)
    }
    for i := 0; i < 5; i++ {
        go writeGo(ch, i+1)
    }
    // keep running
    <-quit
}

func readGo(in <-chan int, idx int) {
    for {
        rwMutex.RLock() // 这样写 得到解锁
        // 🔐 的粒度问题 readLock
        num := <-in
        fmt.Printf("----%d read %d\n", idx, num)
        rwMutex.RUnlock()
    }
}

func writeGo(out chan<- int, idx int) {
    rwMutex.Lock()
    for {
        num := rand.Intn(1000)
        rwMutex.Lock()
        out <- num
        fmt.Printf("%d write %d\n", idx, num)
        time.Sleep(time.Millisecond * 200)
        rwMutex.Unlock()
    }
}

// func readGo(in <-chan int, idx int) {
//     for {
//         rwMutexG.RLock() // 以读模式加锁
//         num := <-in
//         fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
//         rwMutexG.RUnlock() // 以读模式解锁
//     }
// }
//
// func writeGo(out chan<- int, idx int) {
//     for {
//         // 生成随机数
//         num := rand.Intn(1000)
//         rwMutexG.Lock() // 以写模式加锁
//         out <- num
//         fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
//         time.Sleep(time.Millisecond * 300) // 放大实验现象
//         rwMutexG.Unlock()
//     }
// }
