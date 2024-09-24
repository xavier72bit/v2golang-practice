/*
前两个例子里，只返回用户请求的URL看起来有点乏味
我们在这个例子里，为用户返回更多Go语言中http.Request结构体的内容
*/
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:6578", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>%s %s %s</h1>", r.Method, r.URL, r.Proto)
	// 循环打印请求头中的内容
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q <br>", k, v)
	}
	// 打印一下请求主机
	fmt.Fprintf(w, "Host = %q <br>", r.Host)
	// 打印一下远端地址
	fmt.Fprintf(w, "RemoteAddr = %q <br>", r.RemoteAddr)
	// 解析一下Form表单，Go语言允许这样的一个简单的语句结果作为局部的变量
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "解析表单失败: %s <br>", err)
	} else {
		for k, v := range r.Form {
			// 想要看到表单，用这个URL访问: localhost:6578/abc?a=b&c=d&e=f
			fmt.Fprintf(w, "Form[%q] = %q <br>", k, v)
		}
	}
}
