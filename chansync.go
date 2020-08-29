package main

import (
	"fmt"
	"net"
	"os"
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
	// bufferedChan := make(chan int, 9)
	// fmt.Println("cap:", cap(bufferedChan))
	// go func() {
	//     for i := 1; i < 8; i++ {
	//         fmt.Println("in son len() before", len(bufferedChan))
	//         bufferedChan <- i
	//         fmt.Println("--------------son write -----------------------", i)
	//         fmt.Println("in son len() after", len(bufferedChan))
	//     }
	// }()
	// time.Sleep(time.Second)
	// for i := 1; i < 8; i++ {
	//     fmt.Println("in main len() before", len(bufferedChan))
	//     num := <-bufferedChan
	//     fmt.Println("in main len() after", len(bufferedChan))
	//     fmt.Println("main read,", num)
	// }
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		fmt.Println("err during net.Dial", err)
		return
	}
	fmt.Println("connect success to server ~")
	defer conn.Close()
	go func() {
		buffer := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				fmt.Println("err during Stdin.Read err", err)
				continue
			}
			_, err = conn.Write(buffer[:n])
			if err != nil {
				fmt.Println("conn.Write err", err)
				return
			}
		}
	}()

	// get from server
	sWriterBuf := make([]byte, 1024)
	for {
		n, err := conn.Read(sWriterBuf)
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Printf("get from server %s\n", string(sWriterBuf[:n]))
	}
}
