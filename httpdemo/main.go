package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

var BufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}

func main() {
	// 存在两种加Handler的方式
	// HandleFunc() 函数handler
	// Handle() 服务handler e.g. websocket
	http.HandleFunc(
		"/api/hello",
		func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				fmt.Fprint(w, "Err method")
				return
			}

			switch r.Header.Get("Content-Type") {
			case "text/plain":
				// 使用bytes.Buffer, syncPool优化读取big text
				buffer := BufferPool.Get().(*bytes.Buffer)
				buffer.Reset()
				defer BufferPool.Put(buffer)
				_, err := io.Copy(buffer, r.Body)
				defer r.Body.Close()
				if err != nil {
					fmt.Fprint(w, "Err read http body fail")
					return
				}
				fmt.Println(buffer.String())
			case "application/x-www-form-urlencoded":
				r.ParseForm()
				fmt.Println("name: " + r.PostForm.Get("name"))
				fmt.Println("hello: " + r.PostForm.Get("hello"))
			case "application/json":
				// 直接使用jsoniter优化读取json，内部已使用sync.Pool及buf优化
				var b map[string]string
				d := json.NewDecoder(r.Body)
				defer r.Body.Close()
				err := d.Decode(&b)
				if err != nil {
					fmt.Fprint(w, err)
				}
				fmt.Println("name: " + b["name"])
				fmt.Println("hello: " + b["hello"])
			default:
				fmt.Fprint(w, "Err Content-Type")
				return
			}

			fmt.Fprint(w, "Server 127.0.0.1:8888")
		},
	)
	http.ListenAndServe(":8888", nil)
}
