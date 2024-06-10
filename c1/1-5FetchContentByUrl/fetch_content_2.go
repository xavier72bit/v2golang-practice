/*
基于fetch_content_1.go程序进行优化：
使用io.Copy(dst, src)这个函数替代掉fetch_content_1.go中的ioutil.ReadAll(contentByte)来拷贝响应结构体到os.Stdout，避免申请一个缓冲区（例子中的contentByte）。
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// 处理程序参数
	for _, url := range os.Args[1:] {
		// 发起GET请求，获取网络资源
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch %v\n", err)
			os.Exit(1)
		}

		// io.Copy，将响应体的内容拷贝到标准输出上
		/*
			注意啦，这里不能这么写：`_, err := io.Copy(os.Stdout, response.Body)`
			为啥？因为“no new variables on left side of :=”，err已经在上面被定义过咯
		*/
		_, err = io.Copy(os.Stdout, response.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read Content: %v \n", err)
			os.Exit(1)
		}
	}
}
