package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		当数据量巨大的时候，不要像command_line_args_2.go那样采用 `s += sep + s` 的方式进行字符串拼接

		使用join进行拼接：这种用法跟python也很像：
		" ".join(os.args[1:])
	*/

	fmt.Println(strings.Join(os.Args[1:], " "))

	// fmt.Println是可以格式化输出的
	fmt.Println(os.Args[1:])

	// 这个例子将把range的返回值都用上
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
