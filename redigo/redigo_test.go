package main

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "testing"
)

func Test_redigo(t *testing.T) {
    conn, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println(err)
        return
    }
    res, _ := conn.Do("SET", "go", "redigo")
    fmt.Println(res)
}
