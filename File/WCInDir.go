package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

// Need To:
// need to count specific word in a dir and the file ends with txt

// solution:
// 1. find the dir ReadOnly
// 2. find .txt
// 3. open file read line
// 4. split line and save to []string
// 5. iter and count the word
func main() {
    fmt.Println("type the dir")
    var path string
    fmt.Scan(&path)
    file, err := os.OpenFile(path, os.O_RDONLY, os.ModeDir)
    if err != nil {
        fmt.Println("Error happen during open file")
        return
    }
    defer file.Close()
    infos, err := file.Readdir(-1)
    if err != nil {
        fmt.Println("err happened during read dir")
        return
    }
    var dirCount int
    for _, info := range infos {
        // not dir
        if !info.IsDir() {
            if strings.HasSuffix(info.Name(), ".txt") {
                fileCount := wcWordInAFile(path, info, "lol")
                dirCount += fileCount
            }
        }
    }
    fmt.Println("total wcCount in path", path, "is", dirCount)
}

func wcWordInAFile(filePath string, info os.FileInfo, specWord string) int {
    // Mark
    file, err := os.OpenFile(filePath+"/"+info.Name(), os.O_RDONLY, 6)
    if err != nil {
        fmt.Println("err during open file")

    }
    defer file.Close()
    var fileCount int
    reader := bufio.NewReader(file)
    // read the file and count specWord
    for {
        buf, err := reader.ReadBytes('\n')
        fields := strings.Fields(string(buf))
        for _, word := range fields {
            if word == specWord {
                fileCount++
            }
        }
        if err != nil && err == io.EOF {
            fmt.Println("Read File Done")
            return fileCount
        } else if err != nil {
            fmt.Println("ReadBytes err")
        }

    }
}
