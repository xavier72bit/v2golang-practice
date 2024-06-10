/*
os.Args是一个字符串的切片，可以将其类比为Python中的list[str]
它同样支持索引，切片操作，语法跟Python也一样

在command_line_args_1.go中，简单的介绍了一下for关键字，这里将介绍for关键字的另一种用法：

	在字符串、切片等其他数据类型上，可以使用range关键字配合for进行循环。
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		变量声明有以下四种：
		s := ""                [这种只能在函数定义中使用，不能用于包变量]
		var s string
		var s = ""
		var s string = ""
	*/
	s, sep := "", ""

	/*
		每一次循环迭代slice，range将会产生一对值： 索引及其索引处的值，range不管你用不用索引，反正它给你

		在下面这个例子中，我们用不上索引，还不能丢弃它，这样参数不匹配，
		定义一个临时变量用于接收index，也不行，因为GO语言不让出现无用的局部变量，编译都不通过：

			for index, arg := range os.Args[1:] {
				...
			}

	*/

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	// GO语言中标准解决办法是使用空标识符`_`，空标识符可用于在任何语法需要变量名但程序逻辑不需要的时候。
}
