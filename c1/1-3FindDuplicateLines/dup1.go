/*
第一版的需求是：打印标准输入中多次出现的行，以重复次数开头。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		counts : {'Text Line': 1} : {文中的行: 出现次数}

		map 存储了键/值（key/value）的集合，对集合元素，提供常数时间的存、取或测试操作。
			键可以是任意类型，只要其值能用``==`运算符比较，最常见的例子是字符串；
			值则可以是任意类型。


	*/
	counts := make(map[string]int) //使用内置的make函数创建空map，这里map为映射，底层是Hash，跟Python中的dict类似

	// 处理输入，在实际的运行过程中，使用`ctrl`+`d`终止输入
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 下面的语句完全可以简化为：counts[input.Text()]++
		var inputReadLine string
		inputReadLine = input.Text()
		counts[inputReadLine] = counts[inputReadLine] + 1
	}

	// 处理结果输出
	/*
		range在对map进行处理时，跟处理slice不同，它会返回key与其对应的value，这俩我们都需要，就可以直接拿俩变量进行赋值使用。

		完全可以这么写：
			for line, dup := range counts {
				if dup > 1 {
					fmt.Printf("%d\t%s\n", dup, line)
				}
			}

		我偏不，我就用另一种变量定义的方式，这样也方便理解`:=`的作用
	*/
	var line string
	var dup int

	for line, dup = range counts {
		if dup > 1 {
			fmt.Printf("%d\t%s\n", dup, line)
		}
	}

	/*
		这里的Printf也有很多占位符：%d、%t、%c、%s...，Go程序员称之为动词（verb），在Python中，可以用更优雅的f-string来完成
	*/
}
