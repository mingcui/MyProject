/*
 1. 创建一个基于 for 的简单的循环。使其循环 10 次,并且使用 fmt 包打印出计数
器的值。
2. 用 goto 改写 1 的循环。关键字 for 不可使用。
3. 再次改写这个循环,使其遍历一个 array,并将这个 array 打印到屏幕上。
*/
package forLoop

import (
	"fmt"
)

func main() {
	fmt.Printf("%s\n", "打印1到9：")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	var j int = 0
	fmt.Printf("%s\n", "打印0到5:")
Loop:
	fmt.Printf("%d\n", j)
	if j < 5 {
		j++
		goto Loop
	}
	fmt.Printf("%s\n", "打印数组:")
	a := [5]int{6, 7, 8, 9, 10}
	for k := 0; k < 5; k++ {
		fmt.Printf("%d ", a[k])
	}

}
