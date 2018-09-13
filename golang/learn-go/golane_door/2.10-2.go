// 2.10-2
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var s string = "string"
	var i int = 1
	s = strconv.Itoa(i)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}
