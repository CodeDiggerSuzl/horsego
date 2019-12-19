package main

import (
    "fmt"
    "io"
    "os"
)

func main() {
    // copy file
    file, err := os.Open("./File/test.txt")
    if err != nil {
        fmt.Println("open error", err)
        return
    }
    defer file.Close()
    // create file
    copyFile, err := os.Create("./fileCopy.txt")

    if err != nil {
        fmt.Println("error", err)
        return
    }
    defer copyFile.Close()
    // read => to buffer
    buffer := make([]byte, 4096)

    for {
        read, err := file.Read(buffer)
        if err != nil && err == io.EOF {
            fmt.Println("all done")
            return
        }
        copyFile.Write(buffer[:read])
    }
}

// file read and write
// read by bytes and
