/*
第三版的需求是：
之前的两版，都是以“流”的形式读取数据，并根据需要拆分成多个行。
在这个版本中，将一口气把全部输入数据读到内存中，一次分割为多行，然后处理它们。
*/
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)

	if len(os.Args[1:]) < 1 {
		fmt.Println("需要参数: fileName")
	} else {
		for _, fileName := range os.Args[1:] {
			/*
				使用ioutil的ReadFile，将文件读取为一个“Byte slice”，必须把它转换为string，才能用strings.Split分割。
				实现上，`bufio.Scanner`、`ioutil.ReadFile`和`ioutil.WriteFile`，都使用`*os.File`的`Read`和`Write`方法。
				但是，大多数程序员很少需要直接调用那些低级（lower-level）函数。高级（higher-level）函数，像`bufio`和`io/ioutil`包中所提供的那些方法，用起来要容易点。
			*/
			data, err := ioutil.ReadFile(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				count[line]++
			}
		}
	}

	for line, dup := range count {
		if dup > 1 {
			fmt.Printf("%d\t%s\n", dup, line)
		}
	}
}
