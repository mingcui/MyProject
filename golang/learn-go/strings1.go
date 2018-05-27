/*
建立一个 Go 程序打印下面的内容(到 100 个字符):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
*/
package strings1

import (
	"fmt"
)

func main() {
	str := "A"
	for i := 1; i < 9; i++ {
		fmt.Printf("%s\n", str)
		str = str + "A"

	}
}
