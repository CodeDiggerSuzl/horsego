package goroutine

import (
    "fmt"
    "testing"
)

func TestProviderAndConsumer(t *testing.T) {

    buffer := make(chan int, 6)
    go provider(buffer)
    consumer(buffer)
}

// provider 
func provider(out chan<- int) {
    for i := 0; i < 10; i++ {
        out <- i
    }
    close(out)
}

// consumer 
func consumer(in <-chan int) {
    for num := range in {
        fmt.Println("consumer get", num)
    }
}
