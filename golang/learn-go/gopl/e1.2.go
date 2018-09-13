// 修改 	echo	 程序,使其打印每个参数的索引和值,每个一行
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += sep + arg
		fmt.Printf("index is: %s", index)
		fmt.Printf("args is: %s", arg)
		sep = ""
	}
}
