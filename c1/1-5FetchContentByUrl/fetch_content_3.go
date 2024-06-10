/*
在fetch_content_2.go的基础上，让这个程序更加智能一些：
  - 对于那些没有加HTTP协议前缀的，比如`http://`，我们要自动加上协议前缀
  - 每一次发出GET请求后，打印出它的HTTP Status Code
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const HTTP_PREFIX = "http://"

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, HTTP_PREFIX) {
			url = HTTP_PREFIX + url
		}

		response, err := http.Get(url)
		fmt.Printf("Fetch Status Code: %d\n", response.StatusCode)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch Error: %v\n", err)
			os.Exit(1)
		}

		_, err = io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read Content Error: %v\n", err)
			os.Exit(1)
		}
	}
}
