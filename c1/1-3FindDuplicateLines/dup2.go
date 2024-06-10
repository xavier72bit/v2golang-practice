/*
第二版的需求是：在第一版的基础上，增加对打开文件的支持
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	/*
		获取slice长度跟Python是一样的，使用内置的len()函数
		当file不存在时，就是以os.Stdin作为输入流，跟第一版是一样的
	*/
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileArg := range files {
			// 1. 打开文件
			f, err := os.Open(fileArg)
			/*
				要来了！Go语言中的错误处理法！不像Python中的try-except！
				Open函数返回两个值：
					第一个值是被打开的文件，*os.File
					第二个值是内置error类型的值，如果err等于内置值nil(相当于Python中的None)，那么文件被成功打开
			*/
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			}

			// 2. 处理文件流，统计行数
			countLines(f, counts)

			// 3. 最后别忘了关闭资源
			/*
				由于Go中的错误处理机制是逻辑代码的一部分，并不属于异常
				f.Close()的操作不会在try-except-else-finally中处理
				更不会有Python中的with语句处理必然发生的事件
			*/
			f.Close()
		}
	}

	for line, dup := range counts {
		if dup > 1 {
			fmt.Printf("%d\t%s\n", dup, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	/*
		将第一版中的“扫描文件并统计行数”的功能封装为一个方法

		不是哥们，f是一个指针了，那counts这个形参为啥不是指针啊？
		难道它跟Python一样？map是一个可变对象？支持原处修改？
		---
		map作为参数传递给某函数时，该函数接收这个**引用**的一份拷贝
		被调用函数对map底层数据结构的任何修改，调用者函数都可以通过持有的map引用看到
		在此文件的例子中，countLines函数向counts插入的值，也会被main函数看到
	*/
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
