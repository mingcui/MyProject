/*
编写函数,返回其(两个)参数正确的(自然)数字顺序
f(7,2) → 2,7
f(2,7) → 2,7
*/
package main

import (
	"fmt"
)

func getIntOrder(x1, x2 int) (string, string) {
	if x1 > x2 {
		return fmt.Printf("%s and %s\n", string(x2), string(x1))
	} else {
		return string(x1), string(x2)
	}
}

func main() {
	fmt.Println(getIntOrder(7, 2))
	fmt.Println(getIntOrder(2, 7))
}
