// Echo its command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string

	/*
		Python写法：

		for i in len(os.args):
			...
	*/

	/*
		golang中只有for循环
		for initialization; condition; post {
			循环体
		}

		* initialization语句是可选的，在循环开始前执行。initalization如果存在，必须是一条简单语句（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。
		* condition是一个布尔表达式（boolean expression），其值在每次循环迭代开始时计算。如果为true则执行循环体语句。
		* post语句在循环体执行结束后执行，之后再次对condition求值。condition值为false时，循环结束。
	*/
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)

	/*
		for循环的initialization, condition, pot部分每个都可以省略，如果省略initialization和post，分号也可以省略
		一个经典的loop:

			for condition {

			}

		一个死循环，可以靠break和return退出循环

			for {

			}
	*/
}
