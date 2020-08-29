package main

import (
    "fmt"
    "os"
)

// file dir operation
func main() {
    fmt.Println("input a dir to scan")
    var path string
    fmt.Scan(&path)
    file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
    if err != nil {
        fmt.Println("Error happen during open file")
        return
    }
    defer file.Close()
    // read dirs
    readdir, err := file.Readdir(-1)
    if err != nil {
        fmt.Println("Something went wrong during read dir")
        return
    }
    // iter readdir
    for _, fileinfo := range readdir {
        if fileinfo.IsDir() {
            fmt.Println(fileinfo, "is dir")
        } else {
            fmt.Println(fileinfo, "is file")
        }
    }
}
