package lock

import (
    "fmt"
    "math/rand"
    "sync"
    "testing"
    "time"
)

// var cond sync.Cond // 创建全局条件变量
// 
// // 生产者
// func producer(out chan<- int, idx int) {
//     for {
//         cond.L.Lock()              // 1. 先加🔐  加锁成功 执行代码 加锁失败 阻塞  | 条件变量对应互斥锁加锁
//         for len(out) == cap(out) { // 产品区满 等待消费者消费 循环判断 一直加锁 好处是 不满足条件跳过
//             cond.Wait() // 挂起当前协程， 等待条件变量满足，被消费者唤醒
//         }
//         num := rand.Intn(1000) // 产生一个随机数
//         out <- num             // 写入到 channel 中 （生产）
//         fmt.Printf("%dth 生产者，产生数据 %3d, 公共区剩余%d个数据\n", idx, num, len(out))
//         cond.L.Unlock()         // 生产结束，解锁互斥锁
//         cond.Signal()           // 唤醒 阻塞在条件变量上的 消费者
//         time.Sleep(time.Second) // 生产完休息一会，给其他协程执行机会
//     }
// }
// 
// // 消费者
// func consumer(in <-chan int, idx int) {
//     for {
//         cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
//         for len(in) == 0 { // 产品区为空 等待生产者生产
//             cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
//         }
//         num := <-in // 将 channel 中的数据读走 （消费）
//         fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
//         cond.L.Unlock()                    // 消费结束，解锁互斥锁
//         cond.Signal()                      // 唤醒 阻塞的 生产者
//         time.Sleep(time.Millisecond * 500) // 消费完 休息一会，给其他协程执行机会
//     }
// }
// func Test_cond(t *testing.T) {
//     rand.Seed(time.Now().UnixNano()) // 设置随机数种子
//     quit := make(chan bool)          // 创建用于结束通信的 channel
// 
//     product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
//     cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量
// 
//     for i := 0; i < 5; i++ { // 5个消费者
//         go producer(product, i+1)
//     }
//     for i := 0; i < 3; i++ { // 3个生产者
//         go consumer(product, i+1)
//     }
//     <-quit // 主协程阻塞 不结束
// }
var cond sync.Cond

func producer(out chan<- int, index int) {
    // for {
    // 1. lock
    cond.L.Lock()
    // use for to judge
    for len(out) == cap(out) {
        cond.Wait()
    }
    num := rand.Intn(100)
    out <- num
    fmt.Println("P : producer produce ", num, "with ", len(out), "left")
    cond.L.Unlock()
    // tell other go routine
    cond.Signal()
    time.Sleep(time.Second)
    // }
}

// func consumer(in <-chan int, idx int) {
//     for {
//         cond.L.Lock()      // 条件变量对应互斥锁加锁（与生产者是同一个）
//         for len(in) == 0 { // 产品区为空 等待生产者生产
//             cond.Wait() // 挂起当前协程， 等待条件变量满足，被生产者唤醒
//         }
//         num := <-in // 将 channel 中的数据读走 （消费）
//         fmt.Printf("---- %dth 消费者, 消费数据 %3d,公共区剩余%d个数据\n", idx, num, len(in))
//         cond.L.Unlock()                    // 消费结束，解锁互斥锁
//         cond.Signal()                      // 唤醒 阻塞的 生产者
//         time.Sleep(time.Millisecond * 500) // 消费完 休息一会，给其他协程执行机会
//     }
// }
func consumer(in <-chan int, index int) {
    // for {
    cond.L.Lock()
    for len(in) == 0 {
        cond.Wait()
    }
    num := <-in
    fmt.Println("C : consumer consume ", num, "with ", len(in), "left")
    cond.L.Unlock()
    cond.Signal()
    time.Sleep(time.Second)
    // }
}
func Test_Cond(t *testing.T) {
    rand.Seed(time.Now().UnixNano())
    // create the lock of cond.L
    cond.L = new(sync.Mutex)
    quitSignal := make(chan bool)
    // make the buffer for consumer and producer
    buffer := make(chan int, 2)
    for i := 0; i < 5; i++ {
        go consumer(buffer, i+1)
    }
    for i := 0; i < 10; i++ {
        go producer(buffer, i+1)
    }
    <-quitSignal
}

// func Test_cond(t *testing.T) {
//     rand.Seed(time.Now().UnixNano()) // 设置随机数种子
//     quit := make(chan bool)          // 创建用于结束通信的 channel
//
//     product := make(chan int, 3) // 产品区（公共区）使用channel 模拟
//     cond.L = new(sync.Mutex)     // 创建互斥锁和条件变量
//
//     for i := 0; i < 5; i++ { // 5个消费者
//         go producer(product, i+1)
//     }
//     for i := 0; i < 3; i++ { // 3个生产者
//         go consumer(product, i+1)
//     }
//     <-quit // 主协程阻塞 不结束
// }
