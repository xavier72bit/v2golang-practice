/*
在这个文件中，我们在tiny_server_01的基础上，为整个HTTP服务器增加一种“状态”
即统计整个服务器被访问的次数，并可以通过/count查询到总访问次数
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:6578", nil))
}

/*
为啥这里俩函数要加锁呢？
在这些代码的背后，**服务器每一次接收请求的处理是，都会另起一个goroutine**
然而在并发情况下，假如真的有两个请求同一时刻去更新count，那么这个值可能并不会被正确地增加
*/
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "<h1>URL.Path = %q</h1>", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "<h1>Count %d</h1>", count)
	mu.Unlock()
}
