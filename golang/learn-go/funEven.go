// 编写一个函数用于计算一个 float64 类型的 slice 的平均值

package main

import (
	"fmt"
)

func funEven(slic []float64) (avg float64) {
	sum := 0.0
	if len(slic) == 0 {
		avg = 0
	} else {
		for _, i := range slic {
			sum = sum + i
			avg = sum / float64(len(slic))
		}
	}
	return
}

func main() {
	s := []float64{1.2, 2.2, 3.3}
	fmt.Println(funEven(s))
}
