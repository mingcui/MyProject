// ReaderAt 接口使得可以从指定偏移量处开始读取数据
package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("Read at example")
	p := make([]byte, 6)
	n, err := reader.ReadAt(p, 2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s,%d", p, n)
}
