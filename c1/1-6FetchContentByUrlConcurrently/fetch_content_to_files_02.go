/*
在fetch_content_01中，我们并发地获取URL的内容，并统计了整体运行时间

在这里，我们在fetch_content_01的基础上，会将响应程序运行的结果写到一个文件中。
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	fileName := "/Users/xavierwu/tempworkspace/v2golang-practice-1-6.txt"
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil && os.IsNotExist(err) {
		file, err = os.Create(fileName)
		if err != nil {
			fmt.Printf("创建文件时出错: %s \n", err)
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Printf("读取文件时出错: %s \n", err)
		os.Exit(2)
	}

	// 先执行获取
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	// 再写结果
	for range os.Args[1:] {
		_, err := io.WriteString(file, <-ch)
		if err != nil {
			fmt.Printf("向文件写结果时出错: %s \n", err)
			os.Exit(3)
		}
	}

	// 文件内容分隔符，分割不同的运行结果
	_, err = io.WriteString(file, "================================ \n")
	if err != nil {
		fmt.Printf("向文件写结果时出错: %s \n", err)
		os.Exit(3)
	}

	_ = file.Close()

	fmt.Printf("总用时: %.2fs \n", time.Since(start).Seconds())
}

/*
根据URL获取内容，最后统计运行时长
*/
func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("读取%scontent时出错: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
