/*
解决这个叫做 Fizz-Buzz[23] 的问题:
编写一个程序,打印从 1 到 100 的数字。当是三的倍数就打印“Fizz”代替数字,
当是五的倍数就打印 “Buzz” 。当数字同时是三和五的倍数时,打印 “FizzBuzz” 。
*/

package fizzBuzz

import (
	"fmt"
)

func main() {
	// var  arr :=[100]int
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Printf("%s\n", "FizzBuzz")
		case i%3 == 0:
			fmt.Printf("%s\n", "Fizz")
		case i%5 == 0:
			fmt.Printf("%s\n", "Buzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
