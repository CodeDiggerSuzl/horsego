package http

import (
    "net/http"
    "os"
    "testing"
)

func Test_File(t *testing.T) {
    // register the call-back func
    http.HandleFunc("/", fileHandler)
    // bind listen address
    _ = http.ListenAndServe(":8000", nil)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
    OpenFile(r.URL.String(), w)

}

func OpenFile(path string, w http.ResponseWriter) {
    filePath := "./" + path
    file, err := os.Open(filePath)
    HandleErr(err)
    defer file.Close()

    buf := make([]byte, 2048)
    for {
        n, _ := file.Read(buf)
        if n == 0 {
            return
        }
        _, _ = w.Write(buf[:n])
    }
}
