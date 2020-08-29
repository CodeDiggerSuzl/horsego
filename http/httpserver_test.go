package http

import (
    "net/http"
    "testing"
)

func SecondHandler(w http.ResponseWriter, r *http.Request) {
    _, _ = w.Write([]byte("hello again from golang server"))
}
func Test_HttpServer(t *testing.T) {
    // register call back func
    http.HandleFunc("/secondTest", SecondHandler)

    // bind server listen address
    _ = http.ListenAndServe(":8000", nil)
}
