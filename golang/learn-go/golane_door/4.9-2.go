// 4.9-2
package main

import (
	"fmt"
)

func callback() (s string) {
	s = "call me"
	for i := 0; i <= 9; i++ {
		fmt.Println(s)
	}
	return
}

func main() {
	callback()
}
