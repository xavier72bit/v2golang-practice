/*
这个程序将获取程序参数中对应的url，获取其源文本，然后打印出来
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		/*
			http.Get用来创建HTTP GET请求。
			如果获取过程没有出错，那么会在resp这个结构体中得到访问的请求结果
		*/
		resp, err := http.Get(url)
		/*
			别忘了这个经典的GO式错误处理法
		*/
		if err != nil {
			fmt.Fprint(os.Stderr, "Fetch url: %v\n", err)
			os.Exit(1)
		}

		/*
			使用io.ReadAll函数从response中读取到全部内容，将其结果保存在变量contentByte中
			注意随时关闭资源
		*/
		contentByte, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read Content: %v\n", err)
		}
		/*
			值得注意的是，Printf函数会将结果contentByte写出到标准输出流中，不需要我们再自己转换了：
			fmt.Printf("content: \n %s", string(contentByte))
		*/
		fmt.Printf("content: \n %s", contentByte)
	}
}
