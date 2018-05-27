/*
建立一个程序统计字符串里的字符数量:
asSASA ddd dsjkdsjs dk
同时输出这个字符串的字节数。
替换位置 4 开始的三个字符为 “abc”
*/

package strings2

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	var s string = "asSASA ddd dsjkdsjs dk"
	b := []byte(s)
	fmt.Printf("There are : %d charts \n", len(s))
	fmt.Printf("The bytes is about %d \n", utf8.RuneCount(b))
	fmt.Println(strings.Replace(s, "ASA", "abc", 1))
}
