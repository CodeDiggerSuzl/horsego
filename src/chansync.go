package main

import (
    "fmt"
    "time"
)

func main() {
    // ---------------- unbuffered channel --------------
    // ch := make(chan int)
    // go func() {
    //     for i := 0; i < 5; i++ {
    //         fmt.Println("son write", i)
    //         ch <- i
    //     }
    // }()
    // for i := 0; i < 5; i++ {
    //     read := <-ch
    //     fmt.Println("father read", read)
    // }
    // out put:
    // --------
    // son write 0
    // son write 1
    // father read 0
    // father read 1
    // son write 2
    // son write 3
    // father read 2
    // father read 3
    // son write 4
    // father read 4

    // Reason:
    // the son goroutine print and write to chan
    // main goroutine read from chan but the print is an io operation
    // main goroutine couldn't get cpu
    // son goroutine loops and print again
    // ---------------- unbuffered channel --------------

    // ---------------- buffered channel --------------
    bufferedChan := make(chan int, 9)
    fmt.Println("cap:", cap(bufferedChan))
    go func() {
        for i := 1; i < 8; i++ {
            fmt.Println("in son len() before", len(bufferedChan))
            bufferedChan <- i
            fmt.Println("--------------son write -----------------------", i)
            fmt.Println("in son len() after", len(bufferedChan))
        }
    }()
    time.Sleep(time.Second)
    for i := 1; i < 8; i++ {
        fmt.Println("in main len() before", len(bufferedChan))
        num := <-bufferedChan
        fmt.Println("in main len() after", len(bufferedChan))
        fmt.Println("main read,", num)
    }
}
