package lock

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

var cond sync.Cond

func producer(ch chan<- int, idx int) {
    for {
        cond.L.Lock()
        for len(ch) == 3 {
            cond.Wait()

        }
        num := rand.Intn(1000)
        ch <- num
        fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(ch))
        cond.L.Unlock()
        cond.Signal()
        time.Sleep(time.Second)
    }
}
func consumer(in <-chan int, idx int) {
    for {
        cond.L.Lock()
        for len(in) == 3 {
            cond.Wait()
        }
        num := <-in
        fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
        cond.L.Unlock() // 消费结束，解锁互斥锁
        cond.Signal()   // 唤醒 阻塞的 生产者
        time.Sleep(time.Millisecond * 100)

    }
}
func Test_Cond(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    quit := make(chan bool)
    product := make(chan int, 3)
    cond.L = new(sync.Mutex)
    for i := 0; i < 5; i++ { // 5个消费者
        go producer(product, i+1)
    }
    for i := 0; i < 3; i++ { // 3个生产者
        go consumer(product, i+1)
    }

    <-quit
}
