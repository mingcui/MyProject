// 编写计算一个类型是 float64 的 slice 的平均值的代码
package main

import (
	"fmt"
)

func main() {
	s := []float64{1.2, 2.2, 3.3}
	n := float64(len(s))
	sum := 0.0
	for _, i := range s {
		sum = sum + i
	}
	fmt.Println(sum / n)
}
