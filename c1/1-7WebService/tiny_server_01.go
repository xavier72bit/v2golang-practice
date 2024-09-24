/*
在这个文件中，我们将使用Go语言快速写一个HTTP服务器，它将响应用户的请求，并把用户请求的路径在浏览器中打印出来。
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 每个发送到/的请求都会调用一次handler函数
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:6578", nil))
}

// 注意观察这个函数的参数，发送到这个go server的请求，在go里是一个http.Request结构体
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<b>URL.Path = %q\n</b>", r.URL.Path)
}
