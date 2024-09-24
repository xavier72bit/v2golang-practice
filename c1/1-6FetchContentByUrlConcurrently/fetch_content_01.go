/*
在这里，我们浅尝辄止地体会一下Go语言的并发编程，涉及到goroutine和channel

下面的例子fetchall，和前面1-5小节的fetch程序所要做的工作基本一致。
fetchall的特别之处在于它会同时去获取所有的URL，所以这个程序的总执行时间不会超过执行时间最长的那一个任务，前面的fetch程序执行时间则是所有任务执行时间之和。
fetchall程序只会打印获取的内容大小和经过的时间，不会像之前那样打印获取的内容。
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
	/*
		这里的chan数据类型是channel，创建了一个传递string类型参数的channel
	*/
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		/*
			这条语句用来开启goroutine

			goroutine是一种函数的并发执行方式，而channel是用来在goroutine之间进行参数传递。
			main函数本身也运行在一个goroutine中，而go function则表示创建一个新的goroutine，并在这个新的goroutine中执行这个函数。

			循环几次，就会产生几个go routine，也就是说，如果给此程序传递3个URL，那这个程序就会产生4个goroutine(main一个，3个URL参数三个)
		*/
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		/*
			当一个goroutine尝试在一个channel上做send或者receive操作时，这个goroutine会阻塞在调用处，直到另一个goroutine从这个channel里接收或者写入值，这样两个goroutine才会继续执行channel操作之后的逻辑。
			在这个例子中，每一个fetch函数在执行时都会往channel里发送一个值（ch <- expression），主函数负责接收这些值（<-ch）。
		*/
		fmt.Println(<-ch)
	}

	fmt.Printf("总用时: %.2fs \n", time.Since(start).Seconds())
}

/*
根据URL获取内容，并统计运行时长
*/
func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	/*
		又来！Go的经典错误处理方式
	*/
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	/*
		这里的resp.Body不需要，使用io.Copy，把响应的Body内容拷贝到ioutil.Discard输出流中
		ioutil.Discard可以看作是一个数据垃圾桶，类似于/dev/null
	*/
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("读取%scontent时出错: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	/*
		每当请求返回内容时，fetch函数都会往ch这个channel里写入一个字符串，由main函数里的第二个for循环来处理并打印channel里的这个字符串。
	*/

	/*
		向channel发送数据
	*/
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
