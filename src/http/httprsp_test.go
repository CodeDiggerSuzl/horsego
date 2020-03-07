package http

import (
    "net/http"
    "testing"
)

// 浏览器请求服务器的时候 自动调用该函数
func handler(w http.ResponseWriter, r *http.Request) {
    // w 写回给浏览器的数据
    // r 从客户端读到的数据
    w.Write([]byte("Hello from go server"))
}

func Test_HttpResp(t *testing.T) {
    // 注册回调函数,在服务器
    http.HandleFunc("/test", handler)

    //  绑定服务器监听地址 使用默认的 handler
    _ = http.ListenAndServe(":8000", nil)
}
